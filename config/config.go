package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	requiredEnvVars := []string{"GEMINI_API_KEY", "PINECONE_API_KEY", "PINECONE_INDEX_NAME"}
	for _, key := range requiredEnvVars {
	    if os.Getenv(key) == "" {
	        log.Fatalf("Environment variable %s is required", key)
	    }
    }
}
