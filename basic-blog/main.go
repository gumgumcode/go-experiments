package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts []Post

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", ListPosts).Methods("GET")
	r.HandleFunc("/posts", CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", GetPost).Methods("GET")
	r.HandleFunc("/posts/{id}", UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", DeletePost).Methods("DELETE")

	http.Handle("/", r)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
