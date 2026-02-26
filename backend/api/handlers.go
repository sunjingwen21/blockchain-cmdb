package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/blockchain-cmdb/backend/config"
)

// Handler holds all API handlers
type Handler struct {
	config *config.Config
}

// NewHandler creates a new API handler
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{config: cfg}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Service string `json:"service"`
	Mode    string `json:"mode"`
}

// ConfigResponse represents the configuration response (safe to expose)
type ConfigResponse struct {
	Server struct {
		Port string `json:"port"`
		Mode string `json:"mode"`
	} `json:"server"`
	Database struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Name string `json:"name"`
	} `json:"database"`
	Blockchain struct {
		ChainID int    `json:"chain_id"`
		RPCURL  string `json:"rpc_url"`
	} `json:"blockchain"`
}

// HealthCheck handles the health check endpoint
// @Summary Health check
// @Description Returns the health status of the API
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Status:  "ok",
		Version: "0.1.0",
		Service: "blockchain-cmdb-api",
		Mode:    h.config.Server.Mode,
	})
}

// GetConfig handles the configuration endpoint
// @Summary Get configuration
// @Description Returns the current configuration (safe values only)
// @Tags config
// @Produce json
// @Success 200 {object} ConfigResponse
// @Router /api/v1/config [get]
func (h *Handler) GetConfig(c *gin.Context) {
	resp := ConfigResponse{}
	resp.Server.Port = h.config.Server.Port
	resp.Server.Mode = h.config.Server.Mode
	resp.Database.Host = h.config.Database.Host
	resp.Database.Port = h.config.Database.Port
	resp.Database.Name = h.config.Database.DBName
	resp.Blockchain.ChainID = h.config.Blockchain.ChainID
	resp.Blockchain.RPCURL = h.config.Blockchain.RPCURL
	
	c.JSON(http.StatusOK, resp)
}
