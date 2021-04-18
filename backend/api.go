package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/mux"
)

func getArticleList(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	if r.Method == "OPTIONS" {
		return
	}

	cursor, err := client.Query(context.Background(),
		"SELECT title FROM Articles ORDER BY title ASC")

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}

	articleList := []string{}

	for cursor.Next() {
		var article string
		err = cursor.Scan(&article)

		if err != nil {
			fmt.Println(err)
		}

		articleList = append(articleList, article)
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

	result := client.QueryRow(context.Background(),
		"SELECT id, title, subtitle, sidebar, body from Articles WHERE Title=$1", title)

	err := result.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Sidebar,
		&article.Body)

	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
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

	sidebar, _ := json.Marshal(article.Sidebar)

	_, err := client.Exec(context.Background(),
		"UPDATE Articles "+
			"SET Title=$2, Subtitle=$3, Sidebar=$4, Body=$5 "+
			"WHERE ID=$1",
		articleID, article.Title, article.Subtitle, sidebar, article.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createArticle(article Article) int {
	sidebar, _ := json.Marshal(article.Sidebar)

	_, err := client.Exec(context.Background(),
		"INSERT INTO "+
			"Articles(Title, Subtitle, Sidebar, Body) "+
			"VALUES ($1, $2, $3, $4)",
		article.Title, article.Subtitle, sidebar, article.Body)

	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError
	}

	return 0
}
