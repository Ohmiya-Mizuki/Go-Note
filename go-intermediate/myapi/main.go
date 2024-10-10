package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
	"github.com/yourname/reponame/models"
)



func main() {
	comment1 := models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "test comment1",
		CreatedAt: time.Now(),
	}
	
	comment2 := models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "second comment",
		CreatedAt: time.Now(),
	}
	
	article := models.Article{
		ID: 1,
		Title: "first article",
		Contents: "This is the test article.",
		UserName: "saki",
		NiceNum: 1,
		CommentList: []models.Comment{comment1,comment2},
		CreatedAt: time.Now(),
	}
	
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("%s\n",jsonData)
	
	r := mux.NewRouter()
	
	r.HandleFunc("/hello",handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article",handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list",handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice",handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)
	
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
