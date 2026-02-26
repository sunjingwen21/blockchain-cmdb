package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set Gin mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// Initialize router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "ok",
			"version":   "0.1.0",
			"service":   "blockchain-cmdb-api",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Placeholder routes - to be implemented
		v1.GET("/assets", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Assets endpoint - coming soon"})
		})
		
		v1.GET("/blockchain/tx", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Blockchain transactions - coming soon"})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
