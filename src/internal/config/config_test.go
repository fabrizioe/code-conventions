package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test loading default configuration
	config, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Check default values
	if config.Server.Port != 8080 {
		t.Errorf("Expected port 8080, got %d", config.Server.Port)
	}

	if config.Server.ReadTimeout != 30 {
		t.Errorf("Expected read timeout 30, got %d", config.Server.ReadTimeout)
	}

	if config.Logger.Level != "info" {
		t.Errorf("Expected log level 'info', got '%s'", config.Logger.Level)
	}
}

func TestLoadFromEnv(t *testing.T) {
	// Set environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("LOG_LEVEL", "debug")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("LOG_LEVEL")
	}()

	config, err := Load()
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Check that environment variables override defaults
	if config.Server.Port != 9090 {
		t.Errorf("Expected port 9090, got %d", config.Server.Port)
	}

	if config.Logger.Level != "debug" {
		t.Errorf("Expected log level 'debug', got '%s'", config.Logger.Level)
	}
}