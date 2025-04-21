package gemini

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockAPIClient is a mock implementation of APIClient
type MockAPIClient struct {
	Response []byte
	Err      error
}

// SendRequest returns a mock response or error
func (m *MockAPIClient) SendRequest(url, apiKey string, payload []byte) ([]byte, error) {
	return m.Response, m.Err
}

func TestMain(m *testing.M) {
	// Set up global environment variable for testing
	os.Setenv(apiKeyName, "test")
	code := m.Run()
	os.Unsetenv(apiKeyName)
	os.Exit(code)
}

func TestGenerateText(t *testing.T) {
	mockResponse := []byte(`{
		"candidates": [
			{
				"content": {
					"parts": [
						{"text": "This is a mock response text."}
					]
				}
			}
		]
	}`)

	mockClient := newMockClient(mockResponse, nil)
	textGen := GeminiTextGenerator{}

	result, err := textGen.GenerateText(mockClient, "Test prompt")
	assert.NoError(t, err)

	expected := "This is a mock response text."
	assert.Equal(t, expected, result)
}

func TestGenerateText_Error(t *testing.T) {
	mockClient := newMockClient(nil, errors.New("mock error"))
	textGen := GeminiTextGenerator{}

	_, err := textGen.GenerateText(mockClient, "Test prompt")
	assert.Error(t, err, "mock error")
}

// Helper function to create a mock client
func newMockClient(response []byte, err error) *MockAPIClient {
	return &MockAPIClient{Response: response, Err: err}
}
