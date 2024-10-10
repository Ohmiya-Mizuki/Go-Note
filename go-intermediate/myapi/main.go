package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/hello",handlers.HelloHandler)
	r.HandleFunc("/article",handlers.PostArticleHandler)
	r.HandleFunc("/article/list",handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice",handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)
	
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
