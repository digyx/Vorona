package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Article - Database Schema
type Article struct {
	Title    string
	Subtitle string
	Sidebar  map[string]string
	Body     string
}

var ctx context.Context
var client *mongo.Client
var err error

func main() {
	fmt.Println("Starting server...")

	fmt.Println("Connecting to database...")
	// Connect to MongoDB Database
	ctx = context.Background()
	mongoURI := fmt.Sprintf("mongodb://%s:27017", os.Getenv("MONGO_URI"))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to database.")
	}

	fmt.Println("Connected to database.")

	// Initialize router
	r := mux.NewRouter()

	// Handle Endpoints
	r.HandleFunc("/article/{title}", getArticle).Methods("POST", "OPTIONS")
	r.HandleFunc("/healthcheck", healthcheck)

	startServer(r)
}

func startServer(r *mux.Router) {
	// Start server on port 8080
	server := &http.Server{Addr: ":8080", Handler: r}

	// Start the server in a new thread
	go func() {
		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	fmt.Println("Listening on port 8080.")

	// Wait for SIGTERM or SIGINT signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	// Grab context for stopping the server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Actually stop the server
	fmt.Println("Stopping the server...")
	server.Shutdown(ctx)
	fmt.Println("Server has shutdown.")
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	(*w).Header().Set("Access-Control.Allow-Methods", "POST")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	title := mux.Vars(r)["title"]

	var article Article

	filter := bson.M{"title": title}
	collection := client.Database("vorona").Collection("article")
	collection.FindOne(ctx, filter).Decode(&article)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"subtitle": article.Subtitle,
		"sidebar":  article.Sidebar,
		"body":     article.Body,
	})
}
