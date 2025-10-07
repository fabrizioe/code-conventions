package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Handler contains the dependencies for HTTP handlers
type Handler struct {
	logger    *log.Logger
	startTime time.Time
	requests  int64
}

// HelloResponse represents the JSON response for hello endpoints
type HelloResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string        `json:"status"`
	Uptime    time.Duration `json:"uptime"`
	Timestamp time.Time     `json:"timestamp"`
}

// MetricsResponse represents basic metrics
type MetricsResponse struct {
	Requests int64         `json:"requests"`
	Uptime   time.Duration `json:"uptime"`
	Status   string        `json:"status"`
}

// New creates a new Handler instance
func New(logger *log.Logger) *Handler {
	return &Handler{
		logger:    logger,
		startTime: time.Now(),
		requests:  0,
	}
}

// HealthCheck handles health check requests
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.requests++
	h.logger.Printf("Health check requested from %s", r.RemoteAddr)

	response := HealthResponse{
		Status:    "healthy",
		Uptime:    time.Since(h.startTime),
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Hello handles basic hello requests
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	h.requests++
	h.logger.Printf("Hello requested from %s", r.RemoteAddr)

	response := HelloResponse{
		Message:   "Hello, World!",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// HelloWithName handles personalized hello requests
func (h *Handler) HelloWithName(w http.ResponseWriter, r *http.Request) {
	h.requests++
	vars := mux.Vars(r)
	name := vars["name"]

	h.logger.Printf("Hello with name '%s' requested from %s", name, r.RemoteAddr)

	if name == "" {
		http.Error(w, "Name parameter is required", http.StatusBadRequest)
		return
	}

	response := HelloResponse{
		Message:   fmt.Sprintf("Hello, %s!", name),
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Metrics handles metrics requests
func (h *Handler) Metrics(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Metrics requested from %s", r.RemoteAddr)

	response := MetricsResponse{
		Requests: h.requests,
		Uptime:   time.Since(h.startTime),
		Status:   "running",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
