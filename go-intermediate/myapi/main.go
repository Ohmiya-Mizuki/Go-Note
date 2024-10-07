package main

import (
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
	
	
	
	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
