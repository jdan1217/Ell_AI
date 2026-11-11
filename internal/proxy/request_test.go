package proxy

import (
	"bufio"
	"strings"
	"testing"
)

func TestRequest(t *testing.T) {
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
