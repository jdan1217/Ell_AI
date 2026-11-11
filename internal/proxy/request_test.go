package proxy

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestGETRequest(t *testing.T) {
	req := strings.NewReader("GET /test HTTP/1.1\r\nHost: example.com\r\n\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if got.Method != "GET" {
		t.Errorf("Request method parsed as: %s", got.Method)
	}
	if got.URL.Path != "/test" {
		t.Errorf("Expected path /test, got %s", got.URL.Path)
	}
}

func TestPostRequest(t *testing.T) {
	body := `{"username":"admin","password":"test123"}`
	request := fmt.Sprintf("POST /test HTTP/1.1\r\nHost: example.com\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(body), body)
	req := strings.NewReader(request)
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)

	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if got.Method != "POST" {
		t.Errorf("Request method parsed as: %s", got.Method)
	}
	if got.URL.Path != "/test" {
		t.Errorf("Expected path /test, got %s", got.URL.Path)
	}
	bodyBytes, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatalf("Failed to read request body: %v", err)
	}
	if string(bodyBytes) != body {
		t.Errorf("Expected body as %s, got %s instead", body, string(bodyBytes))
	}
}

func TestURLParam(t *testing.T) {
	request := "POST /users?id=123&role=admin HTTP/1.1\r\nHost: example.com\r\nContent-Type: application/json\r\nAuthorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJhZG1pbiI6dHJ1ZX0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c\r\n\r\n"
	req := strings.NewReader(request)
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("Error occured: %s", err)
	}
	if got.URL.Path != "/users" {
		t.Errorf("Expected path = users, instead got %s", got.URL.Path)
	}
	if got.URL.Query().Get("id") != "123" {
		t.Errorf("Expected id = 123, instead got %s", got.URL.Query().Get("id"))
	}
	if got.URL.Query().Get("role") != "admin" {
		t.Errorf("Expected role = admin, instead got %s", got.URL.Query().Get("role"))
	}
}

func TestMalformedRequest(t *testing.T) {
	req := strings.NewReader("GET /test\r\nHost: example.com\r\n\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err == nil {
		t.Fatal("Expected an error for malformed HTTP request")
	}
	if got != nil {
		t.Errorf("Expected Error, got response")
	}
}

func TestRequestHeaders(t *testing.T) {
	req := strings.NewReader("GET /test HTTP/1.1\r\nHost: example.com\r\nAuthorization: Bearer test-token\r\nContent-Type: application/json\r\nUser-Agent: Ell_AI\r\n\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatal("Error occured on request")
	}
	if got.Header.Get("Authorization") != "Bearer test-token" {
		t.Errorf("Expected %s, got %s", "Bearer test-token", got.Header.Get("Authorization"))
	}
	if got.Header.Get("Content-type") != "application/json" {
		t.Errorf("Expected %s, got %s", "application/json", got.Header.Get("Content-type"))
	}
	if got.Header.Get("User-Agent") != "Ell_AI" {
		t.Errorf("Expected %s, got %s", "Ell_AI", got.Header.Get("User-Agent"))
	}
}
