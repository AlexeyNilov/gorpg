package gemini

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	apiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
	apiKeyName = "GOOGLE_GENAI_API_KEY"
)

// APIClient defines an interface for making API requests
type APIClient interface {
	SendRequest(url, apiKey string, payload []byte) ([]byte, error)
}

// DefaultAPIClient is the real implementation of APIClient
type DefaultAPIClient struct{}

// SendRequest makes an HTTP request to the given URL with the payload
func (c *DefaultAPIClient) SendRequest(url, apiKey string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, wrapError("Error creating request", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = "key=" + apiKey

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, wrapError("Error making API request", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response body", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed with status " + resp.Status + ": " + string(body))
	}

	return body, nil
}

type DefaultTextGenerator struct{}

// GenerateText generates text for a given prompt using the Gemini API
func (g *DefaultTextGenerator) GenerateText(client APIClient, prompt string) (string, error) {
	apiKey := getEnvVar(apiKeyName)

	// Prepare request payload
	payload := mustMarshal(map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": prompt}}},
		},
	})

	// Make the API request
	responseBytes, err := client.SendRequest(apiEndpoint, apiKey, payload)
	if err != nil {
		return "", err
	}

	// Extract the text field from the response
	return extractText(responseBytes)
}

// getEnvVar retrieves an environment variable or returns an error if not found
func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

// mustMarshal marshals an object to JSON and returns an error if failure occurs
func mustMarshal(data interface{}) []byte {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	return payload
}

// extractText extracts the text field from the response JSON and returns an error if not found
func extractText(response []byte) (string, error) {
	var parsedResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(response, &parsedResponse); err != nil {
		return "", wrapError("Error parsing response JSON", err)
	}

	if len(parsedResponse.Candidates) > 0 && len(parsedResponse.Candidates[0].Content.Parts) > 0 {
		return parsedResponse.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", errors.New("no text found in response")
}

// wrapError creates a detailed error message
func wrapError(message string, err error) error {
	return errors.New(message + ": " + err.Error())
}
