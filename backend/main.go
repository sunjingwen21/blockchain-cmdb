package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sunjingwen21/blockchain-cmdb/backend/api"
	"github.com/sunjingwen21/blockchain-cmdb/backend/blockchain"
	"github.com/sunjingwen21/blockchain-cmdb/backend/config"
	"github.com/sunjingwen21/blockchain-cmdb/backend/database"
	"github.com/sunjingwen21/blockchain-cmdb/backend/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := database.Init(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	database.DB = db
	defer database.Close(db)

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize blockchain client
	var blockchainClient *blockchain.Client
	if cfg.Blockchain.RPCURL != "" {
		blockchainClient, err = blockchain.NewClient(&cfg.Blockchain)
		if err != nil {
			log.Printf("Warning: Failed to initialize blockchain client: %v", err)
		} else {
			log.Println("Blockchain client initialized successfully")
			defer blockchainClient.Close()
		}
	}

	// Initialize router
	r := gin.Default()

	// Initialize handlers
	handler := api.NewHandler(cfg)
	userHandler := api.NewUserHandler()
	assetHandler := api.NewAssetHandler()
	authHandler := api.NewAuthHandler(cfg.JWT.Secret, cfg.JWT.ExpiresIn)

	// Initialize JWT middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWT.Secret, cfg.JWT.ExpiresIn)

	// Health check endpoint
	r.GET("/health", handler.HealthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Auth routes (public)
		v1.POST("/auth/login", authHandler.Login)
		v1.POST("/auth/register", authHandler.Register)

		// Config routes (public)
		v1.GET("/config", handler.GetConfig)

		// Protected routes (require authentication)
		authorized := v1.Group("")
		authorized.Use(jwtMiddleware.AuthMiddleware())
		{
			// Auth protected routes
			authorized.GET("/auth/me", authHandler.GetCurrentUser)
			authorized.POST("/auth/change-password", authHandler.ChangePassword)

			// User routes
			authorized.GET("/users", userHandler.ListUsers)
			authorized.GET("/users/:id", userHandler.GetUser)
			authorized.POST("/users", middleware.RequireRole("admin"), userHandler.CreateUser)
			authorized.PUT("/users/:id", middleware.RequireRole("admin"), userHandler.UpdateUser)
			authorized.DELETE("/users/:id", middleware.RequireRole("admin"), userHandler.DeleteUser)

			// Asset routes
			authorized.GET("/assets", assetHandler.ListAssets)
			authorized.GET("/assets/:id", assetHandler.GetAsset)
			authorized.POST("/assets", assetHandler.CreateAsset)
			authorized.PUT("/assets/:id", assetHandler.UpdateAsset)
			authorized.DELETE("/assets/:id", assetHandler.DeleteAsset)
			authorized.GET("/assets/:id/history", assetHandler.GetAssetHistory)

			// Blockchain routes
			if blockchainClient != nil {
				authorized.GET("/blockchain/status", func(c *gin.Context) {
					c.JSON(200, gin.H{
						"connected": true,
						"chain_id":  blockchainClient.GetChainID().Int64(),
					})
				})
			}
		}
	}

	// Start server
	log.Printf("Server starting on port %s in %s mode", cfg.Server.Port, cfg.Server.Mode)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
