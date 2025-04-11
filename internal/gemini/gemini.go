package gemini

import (
	"context"
	"os"
	
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var Client *genai.Client 

func InitGemini() error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return err
	}
	Client = client
	return nil
}

func CloseGemini() {
	if Client != nil {
		Client.Close()
	}
}

// Helper function for embeddings
func EmbeddingModel(model string) *genai.EmbeddingModel {
	return Client.EmbeddingModel(model)
}
