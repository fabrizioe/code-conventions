package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestNew(t *testing.T) {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	handler := New(logger)

	if handler == nil {
		t.Fatal("Expected handler to be created, got nil")
	}

	if handler.logger != logger {
		t.Fatal("Expected logger to be set correctly")
	}

	if handler.requests != 0 {
		t.Fatalf("Expected requests to be 0, got %d", handler.requests)
	}
}

func TestHealthCheck(t *testing.T) {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	handler := New(logger)

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.HealthCheck)
	handlerFunc.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Check content type
	expected := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("Expected content type %s, got %s", expected, contentType)
	}

	// Check response body
	var response HealthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("Expected status 'healthy', got '%s'", response.Status)
	}

	// Check that request counter was incremented
	if handler.requests != 1 {
		t.Errorf("Expected requests to be 1, got %d", handler.requests)
	}
}

func TestHello(t *testing.T) {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	handler := New(logger)

	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.Hello)
	handlerFunc.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Check response body
	var response HelloResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedMessage := "Hello, World!"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response.Message)
	}

	if response.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", response.Version)
	}
}

func TestHelloWithName(t *testing.T) {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	handler := New(logger)

	// Test with valid name
	req, err := http.NewRequest("GET", "/hello/Alice", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use mux router to handle path variables
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", handler.HelloWithName).Methods("GET")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Check response body
	var response HelloResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedMessage := "Hello, Alice!"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response.Message)
	}
}

func TestMetrics(t *testing.T) {
	logger := log.New(os.Stdout, "[TEST] ", log.LstdFlags)
	handler := New(logger)

	// Make a few requests to increment counter
	handler.requests = 5

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.Metrics)
	handlerFunc.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	// Check response body
	var response MetricsResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Requests != 5 {
		t.Errorf("Expected requests to be 5, got %d", response.Requests)
	}

	if response.Status != "running" {
		t.Errorf("Expected status 'running', got '%s'", response.Status)
	}
}
