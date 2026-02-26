package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Set test environment variables
	os.Setenv("PORT", "9090")
	os.Setenv("GIN_MODE", "test")
	os.Setenv("DB_HOST", "test-host")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("BLOCKCHAIN_CHAIN_ID", "1337")
	
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("GIN_MODE")
		os.Unsetenv("DB_HOST")
		os.Unsetenv("DB_PORT")
		os.Unsetenv("DB_NAME")
		os.Unsetenv("BLOCKCHAIN_CHAIN_ID")
	}()

	cfg := Load()

	if cfg.Server.Port != "9090" {
		t.Errorf("Expected port to be '9090', got '%s'", cfg.Server.Port)
	}

	if cfg.Server.Mode != "test" {
		t.Errorf("Expected mode to be 'test', got '%s'", cfg.Server.Mode)
	}

	if cfg.Database.Host != "test-host" {
		t.Errorf("Expected DB host to be 'test-host', got '%s'", cfg.Database.Host)
	}

	if cfg.Database.Port != 5433 {
		t.Errorf("Expected DB port to be 5433, got %d", cfg.Database.Port)
	}

	if cfg.Database.DBName != "test_db" {
		t.Errorf("Expected DB name to be 'test_db', got '%s'", cfg.Database.DBName)
	}

	if cfg.Blockchain.ChainID != 1337 {
		t.Errorf("Expected chain ID to be 1337, got %d", cfg.Blockchain.ChainID)
	}
}

func TestLoadDefaults(t *testing.T) {
	cfg := Load()

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected default port to be '8080', got '%s'", cfg.Server.Port)
	}

	if cfg.Server.Mode != "debug" {
		t.Errorf("Expected default mode to be 'debug', got '%s'", cfg.Server.Mode)
	}

	if cfg.Database.Port != 5432 {
		t.Errorf("Expected default DB port to be 5432, got %d", cfg.Database.Port)
	}

	if cfg.JWT.ExpiresIn != 24 {
		t.Errorf("Expected default JWT expires in to be 24, got %d", cfg.JWT.ExpiresIn)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	value := getEnv("TEST_KEY", "default")
	if value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	// Test default value
	value = getEnv("NON_EXISTENT_KEY", "default")
	if value != "default" {
		t.Errorf("Expected 'default', got '%s'", value)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	os.Setenv("TEST_INT", "42")
	defer os.Unsetenv("TEST_INT")

	value := getEnvAsInt("TEST_INT", 0)
	if value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}

	// Test default value
	value = getEnvAsInt("NON_EXISTENT_KEY", 100)
	if value != 100 {
		t.Errorf("Expected 100, got %d", value)
	}

	// Test invalid integer
	os.Setenv("INVALID_INT", "not_a_number")
	defer os.Unsetenv("INVALID_INT")
	value = getEnvAsInt("INVALID_INT", 50)
	if value != 50 {
		t.Errorf("Expected default 50 for invalid int, got %d", value)
	}
}
