package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response APIResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Status != "ok" {
		t.Fatalf("expected status ok, got %s", response.Status)
	}

	if response.Message != "server is alive" {
		t.Fatalf("expected message 'server is alive', got %s", response.Message)
	}
}

func TestHealthHandlerMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestHelloHandlerWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Harsh", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response APIResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedMessage := "hello, Harsh!"
	if response.Message != expectedMessage {
		t.Fatalf("expected message %q, got %q", expectedMessage, response.Message)
	}
}

func TestHelloHandlerDefaultName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()

	helloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response APIResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedMessage := "hello, stranger!"
	if response.Message != expectedMessage {
		t.Fatalf("expected message %q, got %q", expectedMessage, response.Message)
	}
}

func TestEchoHandler(t *testing.T) {
	body := `{"name":"Harsh","role":"DevOps Engineer"}`

	req := httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	echoHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response APIResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Status != "ok" {
		t.Fatalf("expected status ok, got %s", response.Status)
	}

	if response.Message != "echo" {
		t.Fatalf("expected message echo, got %s", response.Message)
	}

	data, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("expected data to be map, got %T", response.Data)
	}

	if data["name"] != "Harsh" {
		t.Fatalf("expected name Harsh, got %v", data["name"])
	}
}

func TestEchoHandlerInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/echo", strings.NewReader(`invalid-json`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	echoHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
