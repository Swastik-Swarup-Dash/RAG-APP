package services

import (
	"context"
	"fmt"
	"rag-app/internal/gemini"

	"github.com/google/generative-ai-go/genai"
)


func GenerateAnswer(question string, contextText string) (string, error) {
	
	prompt := fmt.Sprintf(
		"Use this context: %s\nAnswer this question: %s",
		contextText,
		question,
	)

	model := gemini.Client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(
		context.Background(), 
		genai.Text(prompt),   
	)
	
	if err != nil {
		return "", fmt.Errorf("generation failed: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	
	if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
		return string(textPart), nil
	}
	
	return "", fmt.Errorf("unexpected response format")
}
