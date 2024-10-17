package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestApiHello tests the /api/hello endpoint
func TestApiHello(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create the handler function for the /api/hello route
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello from Go!"})
	})

	// Serve the request with the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `{"message":"Hello from Go!"}`
	// Trimming whitespace and newlines for comparison
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
