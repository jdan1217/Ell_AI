package proxy

import (
	"bufio"
	"errors"
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
	defer got.Body.Close()
	if err != nil {
		t.Fatalf("Failed to read request body: %v", err)
	}
	if string(bodyBytes) != body {
		t.Errorf("Expected body as %s, got %s instead", body, string(bodyBytes))
	}
}

func TestURLParam(t *testing.T) {
	request := "POST /users?id=123&role=admin HTTP/1.1\r\nHost: example.com\r\nContent-Type: application/json\r\n\r\n"
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
		t.Fatalf("parseRequest() error = %v", err)
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

func TestEmptyRequest(t *testing.T) {
	req := strings.NewReader("")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err == nil {
		t.Fatal("Expected an error for empty HTTP request")
	}
	if got != nil {
		t.Errorf("Expected Error, got response")
	}
	if !errors.Is(err, io.EOF) {
		t.Errorf("Expected io.EOF, got %v", err)
	}
}

func TestIncompleteRequest(t *testing.T) {
	req := strings.NewReader("GET /test HTTP/1.1\r\nHost: example.com\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err == nil {
		t.Fatal("Expected an error for incomplete HTTP request")
	}
	if got != nil {
		t.Errorf("Expected nil request, got %v", got)
	}
}

func TestMultipleHeaders(t *testing.T) {
	req := strings.NewReader("GET /test HTTP/1.1\r\nHost: example.com\r\nX-Custom-Header: first\r\nX-Custom-Header: second\r\n\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("parseRequest() error = %v", err)
	}
	values := got.Header.Values("X-Custom-Header")
	if len(values) != 2 {
		t.Fatalf("Expected 2 header values, got %d", len(values))
	}
	if values[0] != "first" {
		t.Errorf("Expected first value, got %s", values[0])
	}
	if values[1] != "second" {
		t.Errorf("Expected second value, got %s", values[1])
	}
}

func TestAbsoluteURL(t *testing.T) {
	req := strings.NewReader("GET http://example.com/test HTTP/1.1\r\nHost: example.com\r\n\r\n")
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("parseRequest() error = %v", err)
	}
	if got.URL.Scheme != "http" {
		t.Errorf("Expected http, got %s", got.URL.Scheme)
	}
	if got.URL.Host != "example.com" {
		t.Errorf("Expected example.com, got %s", got.URL.Host)
	}
	if got.URL.Path != "/test" {
		t.Errorf("Expected /test, got %s", got.URL.Path)
	}
	if got.Method != "GET" {
		t.Errorf("Expected GET, got %s", got.Method)
	}
}

func TestRequestBodyIncomplete(t *testing.T) {
	request := "POST /test HTTP/1.1\r\nHost: example.com\r\nContent-Length: 20\r\n\r\nhello"
	req := strings.NewReader(request)
	reader := bufio.NewReader(req)
	got, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("parseRequest() error = %v", err)
	}

	bodyBytes, err := io.ReadAll(got.Body)
	if err == nil {
		t.Fatal("Expected an error when reading incomplete request body")
	}
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("Expected io.ErrUnexpectedEOF, got %v", err)
	}
	t.Logf("Read %d bytes before error: %s", len(bodyBytes), string(bodyBytes))
}

func TestMultipleRequests(t *testing.T) {
	requests := "GET /first HTTP/1.1\r\nHost: example.com\r\n\r\n" + "GET /second HTTP/1.1\r\nHost: example.com\r\n\r\n"
	req := strings.NewReader(requests)
	reader := bufio.NewReader(req)
	first, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("Failed to parse first request: %v", err)
	}
	second, err := parseRequest(reader)
	if err != nil {
		t.Fatalf("Failed to parse second request: %v", err)
	}

	if first.URL.Path != "/first" {
		t.Errorf("Expected Path /first, got %s", first.URL.Path)
	}
	if second.URL.Path != "/second" {
		t.Errorf("Expected Path /second, got %s", second.URL.Path)
	}
	if first.Method != "GET" {
		t.Errorf("Expected GET, got %s", first.Method)
	}
	if second.Method != "GET" {
		t.Errorf("Expected GET, got %s", second.Method)
	}
	third, err := parseRequest(reader)
	if err == nil {
		t.Fatalf("Expected an error while parsing a third time: %v", err)
	}
	if !errors.Is(err, io.EOF) {
		t.Errorf("Expected io.EOF, got %v", err)
	}
	if third != nil {
		t.Errorf("Expected nil request, got %v", third)
	}
}
