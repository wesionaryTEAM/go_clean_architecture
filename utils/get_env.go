package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

// LoadEnv initially load env
func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
