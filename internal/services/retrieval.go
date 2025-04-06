package services

import (
	"context"
	"fmt"
	"rag-app/internal/db"
	"rag-app/internal/gemini" 
	"github.com/google/generative-ai-go/genai"
)

// For Generating a query with embedding using Gemini
func GenerateQueryEmbedding(query string) ([]float32, error) {
	em := gemini.Client.EmbeddingModel("text-embedding-004")
	resp, err := em.EmbedContent(context.Background(), genai.Text(query))
	if err != nil {
		return nil, err
	}
	return resp.Embedding.Values, nil
}

func RetrieveContext(query string) (string, error) {
	queryEmbedding, err := GenerateQueryEmbedding(query)
	if err != nil {
		return "", fmt.Errorf("embedding generation failed: %w", err)
	}

	sql := `
	SELECT content 
	FROM documents 
	WHERE embedding <-> $1 
	ORDER BY embedding <-> $1 
	LIMIT 1;`

	row := db.DB.QueryRow(context.Background(), sql, queryEmbedding)
	
	var context string
	if err := row.Scan(&context); err != nil {
		return "", fmt.Errorf("failed to retrieve context: %w", err)
	}
	return context, nil 
}
