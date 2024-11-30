package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func handlePutRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if string(body) == "" {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Parse the JSON into the User struct
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Simulate updating the user (e.g., in a database)
	fmt.Printf("Updated user: %+v\n", user)

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "User updated successfully"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/user", handlePutRequest)

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
