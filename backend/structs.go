package main

// Article - Database Schema
type Article struct {
	Title    string
	Subtitle string
	Sidebar  map[string]string
	Body     string
}
