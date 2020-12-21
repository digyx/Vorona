package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDatabase() {
	fmt.Println("Connecting to database...")
	// Connect to MongoDB Database
	ctx = context.Background()
	mongoURI := fmt.Sprintf("mongodb://%s:27017", os.Getenv("MONGO_URI"))
	remote, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to database.")
	}

	client = remote.Database("vorona")

	fmt.Println("Connected to database.")
}

func setupDatabase() {
	fmt.Println("Populating database...")

	files, err := ioutil.ReadDir("./articles")

	if err != nil {
		panic("error:  Could not read articles")
	}

	for _, f := range files {
		title := f.Name()
		subtitle, _ := ioutil.ReadFile(fmt.Sprintf("./articles/%s/subtitle.txt", title))
		body, _ := ioutil.ReadFile(fmt.Sprintf("./articles/%s/body.md", title))
		rawSidebar, _ := ioutil.ReadFile(fmt.Sprintf("./articles/%s/sidebar.txt", title))

		sidebar := map[string]string{}

		for _, entry := range strings.Split(string(rawSidebar), "\n") {
			pair := strings.Split(entry, ":")
			sidebar[pair[0]] = pair[1]
		}

		toAdd := Article{
			Title:    title,
			Subtitle: string(subtitle),
			Sidebar:  sidebar,
			Body:     string(markdown.ToHTML(body, nil, nil)),
		}

		client.Collection("articles").InsertOne(ctx, toAdd)
	}

	fmt.Println("Database populated.")
}
