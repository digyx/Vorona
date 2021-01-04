package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Article - Database Schema
type Article struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Title    string
	Subtitle string
	Sidebar  map[string]string
	Body     string
}

// User - Database Schema
type User struct {
	Email    string
	Password string
}
