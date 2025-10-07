package config

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	Server ServerConfig `yaml:"server"`
	Logger LoggerConfig `yaml:"logger"`
}

// ServerConfig represents HTTP server configuration
type ServerConfig struct {
	Port         int `yaml:"port"`
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
	IdleTimeout  int `yaml:"idle_timeout"`
}

// LoggerConfig represents logging configuration
type LoggerConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// Load loads configuration from file and environment variables
func Load() (*Config, error) {
	// Default configuration
	config := &Config{
		Server: ServerConfig{
			Port:         8080,
			ReadTimeout:  30,
			WriteTimeout: 30,
			IdleTimeout:  120,
		},
		Logger: LoggerConfig{
			Level:  "info",
			Format: "json",
		},
	}

	// Load from config file if it exists
	if err := loadFromFile(config, "config/config.yaml"); err != nil {
		// If file doesn't exist, use defaults (not an error)
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	// Override with environment variables
	loadFromEnv(config)

	return config, nil
}

// loadFromFile loads configuration from a YAML file
func loadFromFile(config *Config, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, config)
}

// loadFromEnv loads configuration from environment variables
func loadFromEnv(config *Config) {
	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Server.Port = p
		}
	}

	if readTimeout := os.Getenv("READ_TIMEOUT"); readTimeout != "" {
		if rt, err := strconv.Atoi(readTimeout); err == nil {
			config.Server.ReadTimeout = rt
		}
	}

	if writeTimeout := os.Getenv("WRITE_TIMEOUT"); writeTimeout != "" {
		if wt, err := strconv.Atoi(writeTimeout); err == nil {
			config.Server.WriteTimeout = wt
		}
	}

	if idleTimeout := os.Getenv("IDLE_TIMEOUT"); idleTimeout != "" {
		if it, err := strconv.Atoi(idleTimeout); err == nil {
			config.Server.IdleTimeout = it
		}
	}

	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Logger.Level = level
	}

	if format := os.Getenv("LOG_FORMAT"); format != "" {
		config.Logger.Format = format
	}
}