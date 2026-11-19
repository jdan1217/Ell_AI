package traffic

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewResponseData(t *testing.T) {
	recorder := httptest.NewRecorder()

	recorder.Header().Set("Content-Type", "application/json")
	recorder.WriteHeader(http.StatusCreated)
	recorder.Write([]byte(`{"status":"created"}`))

	resp := recorder.Result()

	newData, err := NewResponseData(resp)
	if err != nil {
		t.Fatalf("NewResponseData returned an error: %v", err)
	}
	if newData.StatusCode != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, newData.StatusCode)
	}
	if newData.Protocol != "HTTP/1.1" {
		t.Errorf("Expected HTTP/1.1, got %s", newData.Protocol)
	}
}
