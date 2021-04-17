package main

// Article - Database Schema
type Article struct {
	ID       int
	Title    string
	Subtitle string
	Sidebar  map[string]string
	Body     string
}

// User - Database Schema
type User struct {
	ID       int
	Email    string
	Password string
}
