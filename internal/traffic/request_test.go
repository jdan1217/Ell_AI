package traffic

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNewRequestData(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost,
		"http://example.com/api/users",
		strings.NewReader(`{"username":"test"}`),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("X-Custom-Header", "value1")
	req.Header.Add("X-Custom-Header", "value2")
	got, err := NewRequestData(req)
	if err != nil {
		t.Fatalf("NewRequestData returned an error: %v", err)
	}

	if got.Method != http.MethodPost {
		t.Errorf("Expected POST, got %s", got.Method)
	}

	if got.URL != "http://example.com/api/users" {
		t.Errorf("Unexpected URL: %s", got.URL)
	}
	expectedBody := `{"username":"test"}`

	if string(got.Body) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, string(got.Body))
	}
	restoredBody, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("Failed to read restored request body: %v", err)
	}

	if string(restoredBody) != expectedBody {
		t.Errorf("Expected restored body %s, got %s", expectedBody, string(restoredBody))
	}
	contentType := got.Headers["Content-Type"]

	if len(contentType) == 0 || contentType[0] != "application/json" {
		t.Errorf("Expected application/json, got %v", contentType)
	}
	if len(got.Headers["X-Custom-Header"]) != 2 {
		t.Errorf("Expected 2 header values, got %d", len(got.Headers["X-Custom-Header"]))
	}
	req.Header["X-Custom-Header"][0] = "modified"

	if got.Headers["X-Custom-Header"][0] != "value1" {
		t.Error("Modifying original request headers changed captured headers")
	}
}
