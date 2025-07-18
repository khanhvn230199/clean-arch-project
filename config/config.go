package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func Load() *Config {
	// Load .env file if it exists
	godotenv.Load()

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost/dbname?sslmode=disable"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
