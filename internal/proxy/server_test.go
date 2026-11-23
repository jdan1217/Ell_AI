package proxy

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jdan1217/Ell_AI/internal/traffic"
)

func TestForwardRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/test" {
			t.Errorf("Expected /test, got %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Expected Bearer test-token, got %s", r.Header.Get("Authorization"))
		}
	}))
	defer server.Close()

	req, err := http.NewRequest("POST", server.URL+"/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer test-token")
	transport := newTransport()
	defer transport.CloseIdleConnections()
	resp, err := forwardRequest(req, transport)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestForwardResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"status":"created"}`))
	}))
	defer server.Close()

	req, err := http.NewRequest("POST", server.URL+"/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	transport := newTransport()
	defer transport.CloseIdleConnections()
	resp, err := forwardRequest(req, transport)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}
	if resp.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected application/json, got %s", resp.Header.Get("Content-Type"))
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	if string(bodyBytes) != `{"status":"created"}` {
		t.Errorf("Expected {\"status\":\"created\"}, got %s", string(bodyBytes))
	}
}

func TestUpstreamError(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:1/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	transport := newTransport()
	defer transport.CloseIdleConnections()

	resp, err := forwardRequest(req, transport)

	if err == nil {
		if resp != nil {
			resp.Body.Close()
		}
		t.Fatal("Expected an error when upstream server is unavailable")
	}

	if resp != nil {
		t.Error("Expected nil response when upstream request fails")
	}
}

func TestConnectionClose(t *testing.T) {
	client, proxyConn := net.Pipe()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer upstream.Close()

	transport := newTransport()
	history := &traffic.History{}
	defer transport.CloseIdleConnections()

	go handleConnection(proxyConn, transport, history)

	request := fmt.Sprintf(
		"GET %s/test HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n",
		upstream.URL,
		strings.TrimPrefix(upstream.URL, "http://"),
	)
	_, err := client.Write([]byte(request))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	reader := bufio.NewReader(client)
	resp, err := http.ReadResponse(reader, nil)
	if err != nil {
		t.Fatalf("Failed to read response: %v", err)
	}
	defer resp.Body.Close()

	if !resp.Close {
		t.Error("Expected response connection to be closed")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if string(body) != "ok" {
		t.Errorf("Expected response body 'ok', got %q", string(body))
	}
	_, err = client.Read(make([]byte, 1))
	if err != io.EOF {
		t.Errorf("Expected connection to close with EOF, got %v", err)
	}

	client.Close()
}
