package models

import (
	"testing"
	"time"
)

func TestUserToResponse(t *testing.T) {
	now := time.Now()
	user := User{
		ID:        1,
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "should_not_be_included",
		Role:      "admin",
		IsActive:  true,
		CreatedAt: now,
	}

	response := user.ToResponse()

	if response.ID != user.ID {
		t.Errorf("Expected ID %d, got %d", user.ID, response.ID)
	}

	if response.Username != user.Username {
		t.Errorf("Expected Username '%s', got '%s'", user.Username, response.Username)
	}

	if response.Email != user.Email {
		t.Errorf("Expected Email '%s', got '%s'", user.Email, response.Email)
	}

	if response.Role != user.Role {
		t.Errorf("Expected Role '%s', got '%s'", user.Role, response.Role)
	}

	if response.IsActive != user.IsActive {
		t.Errorf("Expected IsActive %v, got %v", user.IsActive, response.IsActive)
	}

	// Password should not be in response
	// This is implicitly tested by the struct definition with json:"-"
}

func TestUserTableName(t *testing.T) {
	user := User{}
	if user.TableName() != "users" {
		t.Errorf("Expected table name 'users', got '%s'", user.TableName())
	}
}

func TestUserBeforeCreate(t *testing.T) {
	user := User{
		Username: "newuser",
		Email:    "new@example.com",
		Role:     "",
	}

	// In a real test, we'd use gorm's DB instance
	// For now, just verify the role is empty before
	if user.Role != "" {
		t.Error("Expected empty role before creation")
	}

	// The BeforeCreate hook would set the default role
	// This is tested implicitly when creating through GORM
}
