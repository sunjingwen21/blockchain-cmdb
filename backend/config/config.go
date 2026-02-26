package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Blockchain BlockchainConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Mode string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret    string
	ExpiresIn int // hours
}

// BlockchainConfig holds blockchain-related configuration
type BlockchainConfig struct {
	RPCURL      string
	ChainID     int
	PrivateKey  string
	ContractAddr string
}

// Load loads configuration from environment variables
func Load() *Config {
	// Load .env file if exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			DBName:   getEnv("DB_NAME", "blockchain_cmdb"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			ExpiresIn: getEnvAsInt("JWT_EXPIRES_IN", 24),
		},
		Blockchain: BlockchainConfig{
			RPCURL:       getEnv("BLOCKCHAIN_RPC_URL", "https://polygon-mumbai.g.alchemy.com/v2/demo"),
			ChainID:      getEnvAsInt("BLOCKCHAIN_CHAIN_ID", 80001),
			PrivateKey:   getEnv("BLOCKCHAIN_PRIVATE_KEY", ""),
			ContractAddr: getEnv("BLOCKCHAIN_CONTRACT_ADDR", ""),
		},
	}
}

// getEnv reads an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt reads an environment variable as integer or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
