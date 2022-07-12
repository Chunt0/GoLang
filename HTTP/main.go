package main

// HTTP server basic jobs:
// Process dynamic request -> incoming user requests, log in to accounts, post images
// Serve statis assets -> serve JS, CSS, images to browsers to create a dynamic experience
// Accept connections -> must listen on a specific port to accept connections

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc()
	http.ListenAndServe(":8080", nil)
}
