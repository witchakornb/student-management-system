package config

import (
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv function
func LoadEnv() error {
	return godotenv.Load()
}

// LoadEnv function with file path
func LoadEnvWithPath(path string) error {
	return godotenv.Load(path)
}

// GetEnv function to get the value of a key from the .env file
func GetEnv(key string) string {
	return os.Getenv(key)
}