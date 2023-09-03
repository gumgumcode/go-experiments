package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Us")
}

func main() {
	// Create a new router using Gorilla Mux.
	r := mux.NewRouter()

	// Define routes and their corresponding handlers.
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/about", AboutHandler).Methods("GET")

	// Create a new HTTP server and specify the router.
	http.Handle("/", r)

	// Start the server on port 8080.
	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", nil)
}
