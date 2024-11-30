package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePutRequest(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid PUT request",
			method:         http.MethodPut,
			body:           `{"id":1,"name":"John Doe","email":"john.doe@example.com"}`,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User updated successfully"}`,
		},
		{
			name:           "Invalid HTTP method",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Invalid request method",
		},
		{
			name:           "Invalid JSON body",
			method:         http.MethodPut,
			body:           `{"id":1,"name":"John Doe",email:"john.doe@example.com"}`, // Missing quotes around email key
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid JSON data",
		},
		{
			name:           "Empty body",
			method:         http.MethodPut,
			body:           "",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to read request bodyf",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new request
			req := httptest.NewRequest(tc.method, "/user", bytes.NewBuffer([]byte(tc.body)))
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder to capture the response
			rr := httptest.NewRecorder()

			// Call the handler
			handlePutRequest(rr, req)

			// Check the status code
			if rr.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, rr.Code)
			}
			// Check the response body
			if strings.ReplaceAll(rr.Body.String(), "\n", "") != tc.expectedBody {
				t.Errorf("expected body %q, got %q", tc.expectedBody, rr.Body.String())
			}
		})
	}
}
