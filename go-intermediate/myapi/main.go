package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)
	
	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article...\n")
	}
	http.HandleFunc("/article", postArticleHandler)
	
	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}
	http.HandleFunc("/articleList", articleListHandler)
	
	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n",articleID)
		io.WriteString(w, resString)
	}
	http.HandleFunc("/article/1", articleDetailHandler)
	
	niceArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}
	http.HandleFunc("/article/nice",niceArticleHandler)
	
	commentHandler := func (w http.ResponseWriter, req *http.Request)  {
		io.WriteString(w,"Posting Comment...\n")	
	}
	http.HandleFunc("/comment", commentHandler)
	
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
