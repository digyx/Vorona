package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func connectToDatabase() {
	fmt.Println("Connecting to database...")
	// Connect to Postgres Database
	remote, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URI"))

	if err != nil {
		fmt.Printf("error: cannot connect to postgres\n%v\n", err)
		os.Exit(1)
	}

	client = remote

	fmt.Println("Connected to database.")
}
