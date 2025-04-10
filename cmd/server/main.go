package main

import (
	"log"
	"rag-app/config"
	"rag-app/internal/db"
	"rag-app/internal/gemini"
	"rag-app/internal/api"
)

func main() {
	config.Load()
	// Initializing my database
	if err := db.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.CloseDB()
	// Initializing Gemini
	if err := gemini.InitGemini(); err != nil {
		log.Fatalf("Gemini initialization failed: %v", err)
	}
	defer gemini.CloseGemini()
	// Starting own server
	router := api.SetupRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
