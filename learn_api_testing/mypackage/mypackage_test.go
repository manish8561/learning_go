package mypackage

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetHandler tests the GetHandler function
func TestGetHandler(t *testing.T) {
	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/api/get", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	GetHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the Content-Type header
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Handler returned wrong Content-Type: got %v want %v", contentType, "application/json")
	}

	// Decode and validate the response body
	var response Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	expectedMessage := "Hello from mypackage GET API"
	if response.Message != expectedMessage {
		t.Errorf("Handler returned unexpected body: got %v want %v", response.Message, expectedMessage)
	}
}

// TestGetHandlerInvalidMethod tests the GetHandler with an invalid HTTP method
func TestGetHandlerInvalidMethod(t *testing.T) {
	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/get", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	GetHandler(rr, req)

	// Check the status code for Method Not Allowed
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	// Check the response body
	expectedBody := "Method not allowed\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}

// TestPostHandler tests the PostHandler function
func TestPostHandler(t *testing.T) {
	// Create a valid POST request body
	requestBody := PostRequest{Name: "John", Age: 30}
	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a new HTTP POST request
	req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	PostHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode and validate the response body
	var response Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
	}

	expectedMessage := "Hello, John! You are 30 years old."
	if response.Message != expectedMessage {
		t.Errorf("Handler returned unexpected body: got %v want %v", response.Message, expectedMessage)
	}
}

// TestPostHandlerInvalidMethod tests the PostHandler with an invalid HTTP method
func TestPostHandlerInvalidMethod(t *testing.T) {
	// Create a new HTTP GET request (invalid for this handler)
	req := httptest.NewRequest(http.MethodGet, "/api/post", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	PostHandler(rr, req)

	// Check the status code for Method Not Allowed
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	// Check the response body
	expectedBody := "Method not allowed\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}

// TestPostHandlerInvalidBody tests the PostHandler with an invalid request body
func TestPostHandlerInvalidBody(t *testing.T) {
	// Create an invalid request body
	invalidBody := []byte(`{"name": "John", "age": "thirty"}`) // "age" should be an integer
	req := httptest.NewRequest(http.MethodPost, "/api/post", bytes.NewBuffer(invalidBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	PostHandler(rr, req)

	// Check the status code for Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	// Check the response body
	expectedBody := "Invalid request body\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}
}
