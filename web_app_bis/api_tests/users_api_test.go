package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserAPI(t *testing.T) {
	// Create a new request to the desired endpoint
	req, err := http.NewRequest("GET", "/api/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Mock server handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your mock server logic goes here
		w.WriteHeader(http.StatusOK)
	})

	// Serve the HTTP request to the mock server
	handler.ServeHTTP(rr, req)

	// Check if the status code is what you expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
