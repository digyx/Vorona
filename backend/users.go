package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func login(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	var dbUser User
	filter := bson.M{"email": user.Email}
	client.Collection("users").FindOne(ctx, filter).Decode(&dbUser)

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil || user.Email == "" {
		fmt.Println(err)
		w.WriteHeader(401)
		return
	}

	w.WriteHeader(200)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// Check if the email is taken
	var checkUser User
	filter := bson.M{"email": user.Email}
	client.Collection("users").FindOne(ctx, filter).Decode(&checkUser)

	if checkUser.Email != "" {
		w.WriteHeader(409)
		return
	}

	// Generate Password Hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	user.Password = string(hash)

	// Add user to Database
	_, err = client.Collection("users").InsertOne(ctx, user)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
