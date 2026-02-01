package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// Imagen 3 API endpoint
	imagenEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/imagen-3.0-generate-001:predict"
)

// ImagenClient handles communication with Google's Imagen API
type ImagenClient struct {
	apiKey     string
	httpClient *http.Client
}

// ImagenRequest represents the API request structure
type ImagenRequest struct {
	Instances  []ImagenInstance  `json:"instances"`
	Parameters ImagenParameters  `json:"parameters"`
}

type ImagenInstance struct {
	Prompt string `json:"prompt"`
}

type ImagenParameters struct {
	SampleCount  int    `json:"sampleCount"`
	AspectRatio  string `json:"aspectRatio,omitempty"`
	OutputFormat string `json:"outputFormat,omitempty"`
}

// ImagenResponse represents the API response
type ImagenResponse struct {
	Predictions []ImagenPrediction `json:"predictions"`
	Error       *ImagenError       `json:"error,omitempty"`
}

type ImagenPrediction struct {
	BytesBase64Encoded string `json:"bytesBase64Encoded"`
	MimeType           string `json:"mimeType"`
}

type ImagenError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// NewImagenClient creates a new Imagen API client
func NewImagenClient(apiKey string) *ImagenClient {
	return &ImagenClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

// Generate creates an image from a text prompt
func (c *ImagenClient) Generate(prompt string) ([]byte, error) {
	return c.GenerateWithOptions(prompt, 1, "1:1")
}

// GenerateWithOptions creates an image with custom parameters
func (c *ImagenClient) GenerateWithOptions(prompt string, count int, aspectRatio string) ([]byte, error) {
	reqBody := ImagenRequest{
		Instances: []ImagenInstance{
			{Prompt: prompt},
		},
		Parameters: ImagenParameters{
			SampleCount: count,
			AspectRatio: aspectRatio,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s?key=%s", imagenEndpoint, c.apiKey)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("api request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ImagenResponse
		if json.Unmarshal(body, &errResp) == nil && errResp.Error != nil {
			return nil, fmt.Errorf("api error [%d]: %s", errResp.Error.Code, errResp.Error.Message)
		}
		return nil, fmt.Errorf("api error [%d]: %s", resp.StatusCode, string(body))
	}

	var imgResp ImagenResponse
	if err := json.Unmarshal(body, &imgResp); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if len(imgResp.Predictions) == 0 {
		return nil, fmt.Errorf("no images returned from API")
	}

	// Decode base64 image
	imageData, err := base64.StdEncoding.DecodeString(imgResp.Predictions[0].BytesBase64Encoded)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	return imageData, nil
}

// Stub mode for testing without API
type StubClient struct{}

func NewStubClient() *StubClient {
	return &StubClient{}
}

func (c *StubClient) Generate(prompt string) ([]byte, error) {
	// Return a small valid PNG (1x1 transparent pixel)
	pngData := []byte{
		0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a,
		0x00, 0x00, 0x00, 0x0d, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x06, 0x00, 0x00, 0x00, 0x1f, 0x15, 0xc4,
		0x89, 0x00, 0x00, 0x00, 0x0a, 0x49, 0x44, 0x41,
		0x54, 0x08, 0xd7, 0x63, 0x00, 0x01, 0x00, 0x00,
		0x05, 0x00, 0x01, 0x0d, 0x0a, 0x2d, 0xb4, 0x00,
		0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae,
		0x42, 0x60, 0x82,
	}
	return pngData, nil
}
