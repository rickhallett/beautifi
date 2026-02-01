package api

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

const (
	// Gemini image generation model
	ImageModel = "gemini-3-pro-image-preview"
)

// GeminiClient handles communication with Google's Gemini API
type GeminiClient struct {
	client *genai.Client
	model  string
}

// NewGeminiClient creates a new Gemini API client
func NewGeminiClient(apiKey string) (*GeminiClient, error) {
	ctx := context.Background()
	
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create genai client: %w", err)
	}

	return &GeminiClient{
		client: client,
		model:  ImageModel,
	}, nil
}

// GenerateImage generates an image from a prompt
func (c *GeminiClient) GenerateImage(ctx context.Context, prompt string) ([]byte, error) {
	// Create content with text prompt
	parts := []*genai.Part{
		{Text: prompt},
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	// Generate content
	result, err := c.client.Models.GenerateContent(ctx, c.model, contents, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	// Extract image from response
	if len(result.Candidates) == 0 {
		return nil, fmt.Errorf("no candidates in response")
	}

	for _, part := range result.Candidates[0].Content.Parts {
		if part.InlineData != nil {
			return part.InlineData.Data, nil
		}
	}

	return nil, fmt.Errorf("no image data in response")
}

// GenerateImages generates multiple images from a prompt
func (c *GeminiClient) GenerateImages(ctx context.Context, prompt string, count int) ([][]byte, error) {
	var images [][]byte
	
	for i := 0; i < count; i++ {
		img, err := c.GenerateImage(ctx, prompt)
		if err != nil {
			return images, fmt.Errorf("failed to generate image %d: %w", i+1, err)
		}
		images = append(images, img)
	}
	
	return images, nil
}

// Close is a no-op for the genai client (it doesn't have a Close method)
func (c *GeminiClient) Close() error {
	// genai.Client doesn't require explicit closing
	return nil
}

// GetAPIKey returns the API key from environment
func GetAPIKey() string {
	return os.Getenv("GEMINI_API_KEY")
}
