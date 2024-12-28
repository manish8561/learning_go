package main

import (
	"log"
	"net/http"

	"example.com/u/demo/mypackage"
)

func main() {

	// Register the handler from the mypackage
	http.HandleFunc("/api/get", mypackage.GetHandler)
	http.HandleFunc("/api/post", mypackage.PostHandler)

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
