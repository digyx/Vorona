package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func getArticleList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	cursor, err := client.Collection("articles").Find(ctx, bson.D{})

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	articleList := []string{}
	for cursor.Next(ctx) {
		var article Article
		cursor.Decode(&article)
		articleList = append(articleList, article.Title)
	}

	json.NewEncoder(w).Encode(articleList)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	title := mux.Vars(r)["title"]

	var article Article

	filter := bson.M{"title": title}
	collection := client.Collection("articles")
	collection.FindOne(ctx, filter).Decode(&article)

	if article.Title == "" {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(article)
}
