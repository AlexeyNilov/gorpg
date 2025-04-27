package textgen

import (
	"github.com/AlexeyNilov/gorpg/textgen/gemini"
)

type MockTextGenerator struct {
	Text string
	Err  error
}

// SendRequest returns a mock response or error
func (m *MockTextGenerator) Generate(prompt string) (string, error) {
	return m.Text, m.Err
}

type TextGenerator interface {
	Generate(prompt string) (string, error)
}

type GenericTextGenerator struct{}

func (g *GenericTextGenerator) Generate(prompt string) (string, error) {
	generator := &gemini.GeminiTextGenerator{}
	text, err := generator.GenerateText(&gemini.DefaultAPIClient{}, prompt)
	return text, err
}
