package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func load() {
    er := godotenv.Load()
    if er != nil {
        log.Fatalf("Error Loading .env file: %v", er)
    }

    requiredEnvVars := []string{"GEMINI_API_KEY","DB_CONN"}
    for _, key := range requiredEnvVars {
        if os.Getenv(key) == "" {
            log.Fatalf("Missing required environment variable: %s", key)
        }
    }
}