package config

import (
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	configs := DefaultConfig()
	if configs.Host != "127.0.0.1" {
		t.Errorf("Expected Host to be 127.0.0.1, got %s", configs.Host)
	}
	if configs.Port != 8080 {
		t.Errorf("Expected Port to be 8080, got %d", configs.Port)
	}
}
