package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context
var client *mongo.Database
var err error

func main() {
	fmt.Println("Starting server...")

	connectToDatabase()

	// Initialize router
	r := mux.NewRouter()

	// Handle Endpoints
	r.HandleFunc("/articles/{title}", getArticle).Methods("GET", "OPTIONS")
	r.HandleFunc("/articles", getArticleList).Methods("GET", "OPTIONS")

	r.HandleFunc("/auth", isLoggedIn).Methods("GET", "OPTIONS")
	r.HandleFunc("/auth", login).Methods("POST", "OPTIONS")
	r.HandleFunc("/auth", logout).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/auth", createAccount).Methods("PUT", "OPTIONS")

	r.HandleFunc("/healthcheck", healthcheck)

	startServer(r)
}

func startServer(r *mux.Router) {
	// Start server on port 8080
	server := &http.Server{Addr: ":8080", Handler: r}

	// Start the server in a new thread
	go func() {
		if os.Getenv("DEV_ENV") == "true" { // Development server uses HTTP
			err := server.ListenAndServe()

			if err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}

			return
		}

		key := "/etc/letsencrypt/live/api.vorona.gg/privkey.pem"
		cert := "/etc/letsencrypt/live/api.vorona.gg/cert.pem"
		err := server.ListenAndServeTLS(cert, key)

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
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
