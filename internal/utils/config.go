package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from a .env file
func LoadConfig(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println("⚠️ No .env file found. Using system environment variables.")
	}
	log.Println("🔧 Environment variables loaded from .env file")
}

// GetEnv fetches the value of an environment variable or returns a default value
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
