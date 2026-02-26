package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/blockchain-cmdb/backend/models"
)

// UserHandler handles user-related requests
type UserHandler struct {
	// In real implementation, this would have a service/repository layer
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"oneof=user admin"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Role     string `json:"role" binding:"omitempty,oneof=user admin"`
	IsActive *bool  `json:"is_active"`
}

// ListUsers returns a list of users
// @Summary List users
// @Description Get all users with pagination
// @Tags users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Mock response - in real implementation, query database
	c.JSON(http.StatusOK, gin.H{
		"data": []models.UserResponse{},
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": 0,
		},
	})
}

// GetUser returns a single user by ID
// @Summary Get user
// @Description Get a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.UserResponse
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Mock response
	c.JSON(http.StatusOK, models.UserResponse{
		ID:       uint(id),
		Username: "demo_user",
		Email:    "demo@example.com",
		Role:     "user",
		IsActive: true,
	})
}

// CreateUser creates a new user
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock creation - in real implementation, save to database
	c.JSON(http.StatusCreated, models.UserResponse{
		ID:       1,
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		IsActive: true,
	})
}

// UpdateUser updates an existing user
// @Summary Update user
// @Description Update a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UpdateUserRequest true "User data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock update
	c.JSON(http.StatusOK, models.UserResponse{
		ID:       uint(id),
		Username: req.Username,
		Email:    req.Email,
		Role:     req.Role,
		IsActive: true,
	})
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Delete a user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	c.Status(http.StatusNoContent)
}
