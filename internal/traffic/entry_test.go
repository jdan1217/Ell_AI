package traffic

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNewTrafficEntry(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost,
		"http://example.com/api",
		strings.NewReader(`{"test":true}`),
	)
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	if err != nil {
		t.Fatal(err)
	}

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{"success":true}`)),
		Proto:      "HTTP/1.1",
	}

	duration := 150 * time.Millisecond
	requestData, err := NewRequestData(req)
	if err != nil {
		t.Fatalf("NewRequestData returned an error: %v", err)
	}
	trafficEntry, err := NewTrafficEntry(requestData, resp, duration)
	if err != nil {
		t.Fatalf("NewTrafficEntry returned an error: %v", err)
	}
	if len(trafficEntry.ID) != 32 {
		t.Errorf("Expected ID length of 32, got %d", len(trafficEntry.ID))
	}

	if trafficEntry.Timestamp.IsZero() {
		t.Errorf("Expected timestamp to be initialized")
	}

	if trafficEntry.Duration != duration {
		t.Errorf("Expected duration %v, got %v", duration, trafficEntry.Duration)
	}

	if trafficEntry.Request.Method != http.MethodPost {
		t.Errorf("Expected method POST, got %s", trafficEntry.Request.Method)
	}

	if trafficEntry.Request.URL != "http://example.com/api" {
		t.Errorf("Expected URL http://example.com/api, got %s", trafficEntry.Request.URL)
	}

	if string(trafficEntry.Request.Body) != `{"test":true}` {
		t.Errorf("Unexpected request body: %s", trafficEntry.Request.Body)
	}

	if trafficEntry.Request.Protocol != "HTTP/1.1" {
		t.Errorf("Expected request protocol HTTP/1.1, got %s", trafficEntry.Request.Protocol)
	}

	if trafficEntry.Response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", trafficEntry.Response.StatusCode)
	}

	if string(trafficEntry.Response.Body) != `{"success":true}` {
		t.Errorf("Unexpected response body: %s", trafficEntry.Response.Body)
	}

	if trafficEntry.Response.Protocol != "HTTP/1.1" {
		t.Errorf("Expected response protocol HTTP/1.1, got %s", trafficEntry.Response.Protocol)
	}

	requestBody, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatalf("Failed to read restored request body: %v", err)
	}

	if string(requestBody) != `{"test":true}` {
		t.Errorf("Unexpected restored request body: %s", requestBody)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read restored response body: %v", err)
	}

	if string(responseBody) != `{"success":true}` {
		t.Errorf("Unexpected restored response body: %s", responseBody)
	}
}
