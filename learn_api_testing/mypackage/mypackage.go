package mypackage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response represents the structure of the API response
type Response struct {
	Message string `json:"message"`
}

// GetHandler handles the GET request
func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a response
	response := Response{Message: "Hello from mypackage GET API"}

	// Set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// PostRequest represents the structure of the POST request body
type PostRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PostHandler handles the POST request
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON body
	var request PostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create a response
	response := Response{
		Message: fmt.Sprintf("Hello, %s! You are %d years old.", request.Name, request.Age),
	}

	// Set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
