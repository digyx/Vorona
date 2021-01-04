package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("Hello World"))

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
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	setupSession(w, r, user.Email)

	w.WriteHeader(http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
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
		w.WriteHeader(http.StatusConflict)
		return
	}

	// Generate Password Hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Password = string(hash)

	// Add user to Database
	_, err = client.Collection("users").InsertOne(ctx, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	setupSession(w, r, user.Email)

	w.WriteHeader(http.StatusOK)
}

func setupSession(w http.ResponseWriter, r *http.Request,
	email string) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = true
	session.Values["email"] = email

	session.Save(r, w)
}

func isLoggedIn(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	session, _ := store.Get(r, "session")

	if session.Values["authenticated"] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if session.Values["authenticated"].(bool) == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
