package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	if r.URL.Query()["format"][0] == "html" {
		htmlBody := markdown.ToHTML([]byte(article.Body), nil, nil)
		article.Body = string(htmlBody)
	}

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	articleID := mux.Vars(r)["id"]

	var article Article
	json.NewDecoder(r.Body).Decode(&article)

	if articleID == "new" {
		statusCode := createArticle(article)

		if statusCode != 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}

	article.ID, _ = primitive.ObjectIDFromHex(articleID)
	filter := bson.M{"_id": article.ID}
	_, err := client.Collection("articles").ReplaceOne(ctx, filter, article)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createArticle(article Article) int {
	var conflictArticle Article
	filter := bson.M{"title": article.Title}
	client.Collection("articles").FindOne(ctx, filter).Decode(&conflictArticle)

	if conflictArticle.Title != "" {
		return http.StatusConflict
	}

	articleInsert := Article{
		Title:    article.Title,
		Subtitle: article.Subtitle,
		Sidebar:  article.Sidebar,
		Body:     article.Body,
	}

	_, err := client.Collection("articles").InsertOne(ctx, articleInsert)

	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError
	}

	return 0
}
