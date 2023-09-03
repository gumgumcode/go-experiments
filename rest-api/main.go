package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts []Post
var nextId = 0

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

func ListPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	json.NewDecoder(r.Body).Decode(&post)
	post.Id = nextId
	nextId++
	posts = append(posts, post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	for _, post := range posts {
		if post.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.NotFound(w, r)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	for i, post := range posts {
		if post.Id == id {
			var updatedPost Post
			json.NewDecoder(r.Body).Decode(&updatedPost)
			posts[i] = updatedPost
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	for i, post := range posts {
		if post.Id == id {
			posts = append(posts[:i], posts[i+1:]...)
		}
	}
}

/*
Request samples:

curl -X POST -H "Content-Type: application/json" -d '{"title":"New Post",
"content":"This is the content of the new post."}' http://localhost:8080/posts

curl http://localhost:8080/posts

curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Post
Title", "content":"Updated post content."}' http://localhost:8080/posts/1

curl -X DELETE http://localhost:8080/posts/1

*/
