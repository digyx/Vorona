package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
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
	err := client.QueryRow(context.Background(), "SELECT * FROM Users WHERE Email=$1", user.Email).Scan(
		&dbUser.ID,
		&dbUser.Email,
		&dbUser.Password)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
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

	// Generate Password Hash
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Password = string(hash)

	// Add user to Database
	_, err = client.Exec(context.Background(),
		"INSERT INTO "+
			"Users (Email, Password) "+
			"VALUES ($1, $2)", user.Email, user.Password)

	if err != nil {
		fmt.Println(err)
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

	if !session.Values["authenticated"].(bool) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func isAdmin(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	session, _ := store.Get(r, "session")
	email := session.Values["email"].(string)
	auth := session.Values["authenticated"].(bool)

	if email != "dtingley@twilit.io" || !auth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
