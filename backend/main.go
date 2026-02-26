package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sunjingwen21/blockchain-cmdb/backend/api"
	"github.com/sunjingwen21/blockchain-cmdb/backend/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize router
	r := gin.Default()

	// Initialize handlers
	handler := api.NewHandler(cfg)
	userHandler := api.NewUserHandler()
	assetHandler := api.NewAssetHandler()

	// Health check endpoint
	r.GET("/health", handler.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Config routes
		v1.GET("/config", handler.GetConfig)

		// User routes
		v1.GET("/users", userHandler.ListUsers)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.POST("/users", userHandler.CreateUser)
		v1.PUT("/users/:id", userHandler.UpdateUser)
		v1.DELETE("/users/:id", userHandler.DeleteUser)

		// Asset routes
		v1.GET("/assets", assetHandler.ListAssets)
		v1.GET("/assets/:id", assetHandler.GetAsset)
		v1.POST("/assets", assetHandler.CreateAsset)
		v1.PUT("/assets/:id", assetHandler.UpdateAsset)
		v1.DELETE("/assets/:id", assetHandler.DeleteAsset)
		v1.GET("/assets/:id/history", assetHandler.GetAssetHistory)
	}

	// Start server
	log.Printf("Server starting on port %s in %s mode", cfg.Server.Port, cfg.Server.Mode)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
