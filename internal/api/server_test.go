package api

import (
	"encoding/json"
	"fmt"
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

func TestGetTrafficDefaultPagination(t *testing.T) {
	history := &traffic.History{}

	for i := 1; i <= 60; i++ {
		history.Add(traffic.TrafficEntry{
			ID: fmt.Sprintf("%d", i),
		})
	}

	server := NewServer(history)

	req := httptest.NewRequest(http.MethodGet, "/api/traffic", nil)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", recorder.Code)
	}

	var entries []traffic.TrafficEntry
	if err := json.NewDecoder(recorder.Body).Decode(&entries); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(entries) != 50 {
		t.Fatalf("Expected 50 entries, got %d", len(entries))
	}
}

func TestGetTrafficCustomLimit(t *testing.T) {
	history := &traffic.History{}

	for i := 1; i <= 10; i++ {
		history.Add(traffic.TrafficEntry{
			ID: fmt.Sprintf("%d", i),
		})
	}

	server := NewServer(history)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/traffic?limit=3",
		nil,
	)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", recorder.Code)
	}

	var entries []traffic.TrafficEntry
	if err := json.NewDecoder(recorder.Body).Decode(&entries); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	if entries[0].ID != "1" || entries[2].ID != "3" {
		t.Fatalf("Unexpected entries returned: %+v", entries)
	}
}

func TestGetTrafficCustomOffset(t *testing.T) {
	history := &traffic.History{}

	for i := 1; i <= 10; i++ {
		history.Add(traffic.TrafficEntry{
			ID: fmt.Sprintf("%d", i),
		})
	}

	server := NewServer(history)

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/traffic?limit=3&offset=4",
		nil,
	)
	recorder := httptest.NewRecorder()

	server.Handler().ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", recorder.Code)
	}

	var entries []traffic.TrafficEntry
	if err := json.NewDecoder(recorder.Body).Decode(&entries); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(entries) != 3 {
		t.Fatalf("Expected 3 entries, got %d", len(entries))
	}

	if entries[0].ID != "5" || entries[1].ID != "6" || entries[2].ID != "7" {
		t.Fatalf("Unexpected entries returned: %+v", entries)
	}
}

func TestGetTrafficInvalidPagination(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "invalid limit",
			url:  "/api/traffic?limit=abc",
		},
		{
			name: "zero limit",
			url:  "/api/traffic?limit=0",
		},
		{
			name: "negative limit",
			url:  "/api/traffic?limit=-5",
		},
		{
			name: "negative offset",
			url:  "/api/traffic?offset=-5",
		},
		{
			name: "invalid offset",
			url:  "/api/traffic?offset=abc",
		},
		{
			name: "limit too large",
			url:  "/api/traffic?limit=501",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			history := &traffic.History{}
			server := NewServer(history)

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			recorder := httptest.NewRecorder()

			server.Handler().ServeHTTP(recorder, req)

			if recorder.Code != http.StatusBadRequest {
				t.Fatalf(
					"Expected status 400, got %d",
					recorder.Code,
				)
			}
		})
	}
}
