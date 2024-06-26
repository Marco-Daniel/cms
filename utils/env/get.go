package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key, fallback string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
