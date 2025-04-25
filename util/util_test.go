package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTemplate(t *testing.T) {
	// Define a sample template string
	templateStr := "Hello, {{.Name}}! Welcome to {{.Place}}."

	// Define the data to be used in the template
	data := struct {
		Name  string
		Place string
	}{
		Name:  "Alice",
		Place: "Wonderland",
	}

	// Call the ParseTemplate function
	result := ParseTemplate(templateStr, data)

	// Expected output
	expected := "Hello, Alice! Welcome to Wonderland."

	// Assert that the result matches the expected output
	assert.Equal(t, expected, result)
}

func TestExtractName(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `Name: John Doe
Description: Test description`,
			expected: "John Doe",
		},
		{
			input: `Name: *Jane Smith*
Description: Another test description`,
			expected: "Jane Smith",
		},
		{
			input: `Description: Test description
Name: Alice Wonderland`,
			expected: "Alice Wonderland",
		},
		{
			input: `Description: Test description without a name field`,
			expected: "",
		},
		{
			input: ``,
			expected: "",
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ExtractName(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestExtractDescription(t *testing.T) {
	// Test cases
	tests := []struct {
		input    string
		expected string
	}{
		{
			input: `Name: John Doe
Description: Test description`,
			expected: "Test description",
		},
		{
			input: `Description: Another test description
Name: Jane Smith`,
			expected: "Another test description",
		},
		{
			input: `Name: Alice
Description: Multiline description
continues here`,
			expected: "Multiline description\ncontinues here",
		},
		{
			input: `Name: Bob
SomeOtherField: Some data`,
			expected: "",
		},
		{
			input: ``,
			expected: "",
		},
	}

	// Run test cases
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := ExtractDescription(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
