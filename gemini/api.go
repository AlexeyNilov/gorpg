package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	apiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
)

// GenerateText generates text for a given prompt using the Gemini API
func GenerateText(prompt string) string {
	apiKey := getEnvVar("GOOGLE_GENAI_API_KEY")

	// Prepare request payload
	payload := mustMarshal(map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": prompt}}},
		},
	})

	// Make the API request
	responseBytes := mustSendRequest(apiEndpoint, apiKey, payload)

	// Extract the text field from the response
	return mustExtractText(responseBytes)
}

// getEnvVar retrieves an environment variable or exits if not found
func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

// mustMarshal marshals an object to JSON or exits on failure
func mustMarshal(data interface{}) []byte {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	return payload
}

// mustSendRequest sends an HTTP request and returns the response bytes or exits on failure
func mustSendRequest(url, apiKey string, payload []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = fmt.Sprintf("key=%s", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making API request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status %d: %s", resp.StatusCode, body)
	}

	return body
}

// mustExtractText extracts the text field from the response JSON or exits on failure
func mustExtractText(response []byte) string {
	var parsedResponse struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	err := json.Unmarshal(response, &parsedResponse)
	if err != nil {
		log.Fatalf("Error parsing response JSON: %v", err)
	}

	if len(parsedResponse.Candidates) > 0 && len(parsedResponse.Candidates[0].Content.Parts) > 0 {
		return parsedResponse.Candidates[0].Content.Parts[0].Text
	}

	log.Fatal("No text found in response")
	return ""
}
