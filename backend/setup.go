package main

import (
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func connectToDatabase() {
	fmt.Println("Connecting to database...")
	// Connect to Postgres Database
	config, _ := pgx.ParseConnectionString(os.Getenv("POSTGRES_URI"))
	remote, err := pgx.Connect(config)

	if err != nil {
		fmt.Printf("error: cannot connect to postgres\n%v\n", err)
		os.Exit(1)
	}

	client = remote

	fmt.Println("Connected to database.")
}
