package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jdan1217/Ell_AI/internal/traffic"
)

func TestGetTrafficStatusCode(t *testing.T) {
	history := &traffic.History{}
	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestGetTrafficContentType(t *testing.T) {
	history := &traffic.History{}
	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	contentType := recorder.Header().Get("Content-Type")

	if contentType != "application/json" {
		t.Errorf("Expected application/json, got %s", contentType)
	}
}

func TestGetTrafficEntries(t *testing.T) {
	history := &traffic.History{}

	history.Add(traffic.TrafficEntry{
		ID: "test-123",
	})

	history.Add(traffic.TrafficEntry{
		ID: "test-456",
	})

	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	var entries []traffic.TrafficEntry

	err := json.NewDecoder(recorder.Body).Decode(&entries)
	if err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	if len(entries) != 2 {
		t.Fatalf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].ID != "test-123" {
		t.Errorf("Expected first entry ID test-123, got %s", entries[0].ID)
	}

	if entries[1].ID != "test-456" {
		t.Errorf("Expected second entry ID test-456, got %s", entries[1].ID)
	}
}

func TestGetTrafficByID(t *testing.T) {
	history := &traffic.History{}
	history.Add(traffic.TrafficEntry{ID: "test-123"})

	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic/test-123", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", recorder.Code)
	}

	var entry traffic.TrafficEntry
	if err := json.NewDecoder(recorder.Body).Decode(&entry); err != nil {
		t.Fatal(err)
	}

	if entry.ID != "test-123" {
		t.Errorf("Expected ID test-123, got %s", entry.ID)
	}
}

func TestGetTrafficByIDNotFound(t *testing.T) {
	history := &traffic.History{}
	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic/missing", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", recorder.Code)
	}
}
