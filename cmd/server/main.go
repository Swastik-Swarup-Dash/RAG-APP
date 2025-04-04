package main

import (
    "log"
    "rag-appp/config"
)

func main() {
    config.load()

    if er : db.InitDB(); er != nil {
        log.Fatalf("Database connection error: %v", er)
    }
    defer db.CloseDB()

    if er: gemini.InitGemini(); er!=nil {
        log.Fatalf("Gemini initialization error: %v", er)
    }
    defer gemini.CloseGemini()
    
    router := api.SetupRouter()
    if er := router.Run(":8000"); er!=nil{
      log.Fatalf("Server run error: %v", er)
    }
}