package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// APIResponse is the standard JSON shape every endpoint returns.
// json tags control the exact key names in the response.
type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"` // omitempty = skip key if nil
	Timestamp string      `json:"timestamp"`
}

// respond is a helper so we don't repeat header + encode logic in every handler.
func respond(w http.ResponseWriter, statusCode int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// GET /health
// Kubernetes uses this as a liveness probe.
// If this returns 200, the pod is considered healthy.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	respond(w, http.StatusOK, APIResponse{
		Status:    "ok",
		Message:   "server is alive",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// GET /hello?name=Harsh
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	respond(w, http.StatusOK, APIResponse{
		Status:    "ok",
		Message:   "hello, " + name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// POST /echo
// Reads whatever JSON you send, echoes it back inside "data".
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad request: invalid json", http.StatusBadRequest)
		return
	}
	respond(w, http.StatusOK, APIResponse{
		Status:    "ok",
		Message:   "echo",
		Data:      body,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func main() {
	// PORT env var lets us override the port — useful in different environments.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// ServeMux is Go's built-in HTTP router.
	// In a real service you'd use chi or gin, but stdlib is fine for learning.
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/echo", echoHandler)

	log.Printf("server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server crashed: %v", err)
	}
}
