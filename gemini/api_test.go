package gemini

import (
	"errors"
	"os"
	"testing"
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

	result, err := GenerateText(mockClient, "Test prompt")
	assertNoError(t, err)

	expected := "This is a mock response text."
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestGenerateText_Error(t *testing.T) {
	mockClient := newMockClient(nil, errors.New("mock error"))

	_, err := GenerateText(mockClient, "Test prompt")
	assertError(t, err, "mock error")
}

// Helper function to create a mock client
func newMockClient(response []byte, err error) *MockAPIClient {
	return &MockAPIClient{Response: response, Err: err}
}

// Helper function to assert no error
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// Helper function to assert a specific error message
func assertError(t *testing.T, err error, expectedMessage string) {
	t.Helper()
	if err == nil {
		t.Fatalf("Expected error, but got none")
	}
	if err.Error() != expectedMessage {
		t.Errorf("Expected error %q, but got %q", expectedMessage, err.Error())
	}
}
