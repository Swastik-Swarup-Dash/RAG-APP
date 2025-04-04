package services

import (
	"context"
	"rag-app/internal/db"
	"rag-app/internal/gemini"

	"github.com/google/generative-ai-go/genai"
)

func ProcessDocument(content string) error {
	// Using an EmbeddingModel instead of GenerativeModel
	em := gemini.Client.EmbeddingModel("text-embedding-004") 
	
	resp, err := em.EmbedContent(context.Background(), genai.Text(content))
	if err != nil {
		return err
	}

	// Use exported DB function
	return db.StoreEmbedding(content, resp.Embedding.Values)
}
