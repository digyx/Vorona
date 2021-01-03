package main

import (
	"context"
	"fmt"
	"os"

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
