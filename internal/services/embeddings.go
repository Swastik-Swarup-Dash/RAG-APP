package services

import (
	"context"
	"log"
	"rag-app/internal/db"
	"rag-app/internal/gemini"

	"github.com/google/generative-ai-go/genai"
)

func ProcessDocument(content string) error {
	log.Println("Starting document processing...")

	// Generate embeddings using Gemini
	em := gemini.Client.EmbeddingModel("text-embedding-004")
	resp, err := em.EmbedContent(context.Background(), genai.Text(content))
	if err != nil {
		log.Printf("Error generating embeddings: %v", err)
		return err
	}

	log.Println("Embeddings generated successfully.")

	// Store embeddings in the database
	err = db.StoreEmbedding(content, resp.Embedding.Values)
	if err != nil {
		log.Printf("Error storing embeddings in DB: %v", err)
		return err
	}

	log.Println("Document stored successfully.")
	return nil
}
