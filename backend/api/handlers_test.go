package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/blockchain-cmdb/backend/config"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

func TestHealthCheck(t *testing.T) {
	r := setupTestRouter()
	cfg := &config.Config{
		Server: config.ServerConfig{
			Mode: "test",
		},
	}
	handler := NewHandler(cfg)

	r.GET("/health", handler.HealthCheck)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response HealthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Status != "ok" {
		t.Errorf("Expected status 'ok', got '%s'", response.Status)
	}

	if response.Version != "0.1.0" {
		t.Errorf("Expected version '0.1.0', got '%s'", response.Version)
	}

	if response.Service != "blockchain-cmdb-api" {
		t.Errorf("Expected service 'blockchain-cmdb-api', got '%s'", response.Service)
	}
}

func TestGetConfig(t *testing.T) {
	r := setupTestRouter()
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: "8080",
			Mode: "test",
		},
		Database: config.DatabaseConfig{
			Host: "localhost",
			Port: 5432,
			DBName: "test_db",
		},
		Blockchain: config.BlockchainConfig{
			ChainID: 1337,
			RPCURL:  "http://localhost:8545",
		},
	}
	handler := NewHandler(cfg)

	r.GET("/config", handler.GetConfig)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/config", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response ConfigResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Server.Port != "8080" {
		t.Errorf("Expected port '8080', got '%s'", response.Server.Port)
	}

	if response.Database.Port != 5432 {
		t.Errorf("Expected database port 5432, got %d", response.Database.Port)
	}
}

func TestUserHandlerCreateUser(t *testing.T) {
	r := setupTestRouter()
	handler := NewUserHandler()

	r.POST("/users", handler.CreateUser)

	requestBody := CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}

	body, _ := json.Marshal(requestBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["username"] != "testuser" {
		t.Errorf("Expected username 'testuser', got '%v'", response["username"])
	}
}

func TestUserHandlerCreateUserInvalid(t *testing.T) {
	r := setupTestRouter()
	handler := NewUserHandler()

	r.POST("/users", handler.CreateUser)

	// Invalid request - missing required fields
	invalidBody := map[string]string{
		"username": "ab", // too short
	}

	body, _ := json.Marshal(invalidBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d for invalid request, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestAssetHandlerCreateAsset(t *testing.T) {
	r := setupTestRouter()
	handler := NewAssetHandler()

	r.POST("/assets", handler.CreateAsset)

	requestBody := CreateAssetRequest{
		Name: "Test Server",
		Type: "hardware",
		Cost: 1000.00,
	}

	body, _ := json.Marshal(requestBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/assets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["name"] != "Test Server" {
		t.Errorf("Expected name 'Test Server', got '%v'", response["name"])
	}
}

func TestAssetHandlerCreateAssetInvalid(t *testing.T) {
	r := setupTestRouter()
	handler := NewAssetHandler()

	r.POST("/assets", handler.CreateAsset)

	// Invalid request - invalid asset type
	invalidBody := map[string]interface{}{
		"name": "Test Asset",
		"type": "invalid_type",
	}

	body, _ := json.Marshal(invalidBody)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/assets", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d for invalid request, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestUserHandlerGetUser(t *testing.T) {
	r := setupTestRouter()
	handler := NewUserHandler()

	r.GET("/users/:id", handler.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response["id"] != float64(1) {
		t.Errorf("Expected id 1, got %v", response["id"])
	}
}

func TestUserHandlerGetUserInvalidID(t *testing.T) {
	r := setupTestRouter()
	handler := NewUserHandler()

	r.GET("/users/:id", handler.GetUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/invalid", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d for invalid ID, got %d", http.StatusBadRequest, w.Code)
	}
}
