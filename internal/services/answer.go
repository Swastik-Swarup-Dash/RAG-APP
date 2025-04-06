package services

import (
	"context"
	"fmt"
	"rag-app/internal/gemini"

	"github.com/google/generative-ai-go/genai"
)

// GenerateAnswer combines retrieved context and user question
func GenerateAnswer(question string, contextText string) (string, error) {
	// Use proper parameter name to avoid shadowing context package
	prompt := fmt.Sprintf(
		"Use this context: %s\nAnswer this question: %s",
		contextText,
		question,
	)

	model := gemini.Client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(
		context.Background(), // Proper context package usage
		genai.Text(prompt),   // Correct text part creation
	)
	
	if err != nil {
		return "", fmt.Errorf("generation failed: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	// Convert genai.Part to string safely
	if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return string(textPart), nil
	}
	
	return "", fmt.Errorf("unexpected response format")
}
