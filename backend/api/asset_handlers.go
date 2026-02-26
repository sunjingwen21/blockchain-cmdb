package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/blockchain-cmdb/backend/models"
)

// AssetHandler handles asset-related requests
type AssetHandler struct {
	// In real implementation, this would have a service/repository layer
}

// NewAssetHandler creates a new asset handler
func NewAssetHandler() *AssetHandler {
	return &AssetHandler{}
}

// CreateAssetRequest represents the request body for creating an asset
type CreateAssetRequest struct {
	Name           string             `json:"name" binding:"required,max=200"`
	Description    string             `json:"description"`
	Type           models.AssetType   `json:"type" binding:"required,oneof=hardware software license network other"`
	Location       string             `json:"location"`
	PurchaseDate   string             `json:"purchase_date" binding:"omitempty,datetime=2006-01-02"`
	WarrantyExpiry string             `json:"warranty_expiry" binding:"omitempty,datetime=2006-01-02"`
	Cost           float64            `json:"cost" binding:"min=0"`
	SerialNumber   string             `json:"serial_number"`
	Metadata       models.JSONMap     `json:"metadata"`
}

// UpdateAssetRequest represents the request body for updating an asset
type UpdateAssetRequest struct {
	Name           string             `json:"name" binding:"omitempty,max=200"`
	Description    string             `json:"description"`
	Status         models.AssetStatus `json:"status" binding:"omitempty,oneof=active inactive maintenance retired"`
	Location       string             `json:"location"`
	WarrantyExpiry string             `json:"warranty_expiry" binding:"omitempty,datetime=2006-01-02"`
	Cost           float64            `json:"cost" binding:"omitempty,min=0"`
	Metadata       models.JSONMap     `json:"metadata"`
}

// ListAssets returns a list of assets
// @Summary List assets
// @Description Get all assets with pagination and filtering
// @Tags assets
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param type query string false "Asset type filter"
// @Param status query string false "Status filter"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/assets [get]
func (h *AssetHandler) ListAssets(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	assetType := c.Query("type")
	status := c.Query("status")

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Mock response - in real implementation, query database
	c.JSON(http.StatusOK, gin.H{
		"data": []models.AssetResponse{},
		"pagination": gin.H{
			"page":   page,
			"limit":  limit,
			"total":  0,
			"type":   assetType,
			"status": status,
		},
	})
}

// GetAsset returns a single asset by ID
// @Summary Get asset
// @Description Get an asset by ID
// @Tags assets
// @Produce json
// @Param id path int true "Asset ID"
// @Success 200 {object} models.AssetResponse
// @Failure 404 {object} map[string]string
// @Router /api/v1/assets/{id} [get]
func (h *AssetHandler) GetAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset id"})
		return
	}

	// Mock response
	c.JSON(http.StatusOK, models.AssetResponse{
		ID:       uint(id),
		AssetID:  "AST-20240226-00001",
		Name:     "Demo Asset",
		Type:     "hardware",
		Status:   "active",
		Location: "Server Room A",
	})
}

// CreateAsset creates a new asset
// @Summary Create asset
// @Description Create a new asset
// @Tags assets
// @Accept json
// @Produce json
// @Param asset body CreateAssetRequest true "Asset data"
// @Success 201 {object} models.AssetResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/assets [post]
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var req CreateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock creation - in real implementation, save to database and blockchain
	c.JSON(http.StatusCreated, models.AssetResponse{
		ID:          1,
		AssetID:     "AST-20240226-00001",
		Name:        req.Name,
		Description: req.Description,
		Type:        string(req.Type),
		Status:      "active",
		Location:    req.Location,
		Cost:        req.Cost,
	})
}

// UpdateAsset updates an existing asset
// @Summary Update asset
// @Description Update an asset by ID
// @Tags assets
// @Accept json
// @Produce json
// @Param id path int true "Asset ID"
// @Param asset body UpdateAssetRequest true "Asset data"
// @Success 200 {object} models.AssetResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/assets/{id} [put]
func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset id"})
		return
	}

	var req UpdateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mock update
	c.JSON(http.StatusOK, models.AssetResponse{
		ID:     uint(id),
		Name:   req.Name,
		Status: string(req.Status),
	})
}

// DeleteAsset deletes an asset
// @Summary Delete asset
// @Description Delete an asset by ID
// @Tags assets
// @Produce json
// @Param id path int true "Asset ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /api/v1/assets/{id} [delete]
func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset id"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAssetHistory returns the blockchain history for an asset
// @Summary Get asset history
// @Description Get blockchain transaction history for an asset
// @Tags assets
// @Produce json
// @Param id path int true "Asset ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/assets/{id}/history [get]
func (h *AssetHandler) GetAssetHistory(c *gin.Context) {
	_, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid asset id"})
		return
	}

	// Mock blockchain history
	c.JSON(http.StatusOK, gin.H{
		"asset_id": c.Param("id"),
		"transactions": []gin.H{
			{
				"tx_hash":   "0x...",
				"event":     "created",
				"timestamp": "2024-02-26T10:00:00Z",
				"data":      gin.H{},
			},
		},
	})
}
