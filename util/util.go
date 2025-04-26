package util

import (
	"bytes"
	"regexp"
	"strings"
	"text/template"
)

func ParseTemplate(templateStr string, data any) string {
	// Parse the template
	tpl, err := template.New("New").Parse(templateStr)
	if err != nil {
		panic(err)
	}

	// Use a bytes.Buffer to capture the output
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func ExtractName(input string) string {
	// Define the regular expression to extract the description
	re := regexp.MustCompile(`(?m)^Name:\s*(.+)$`)

	// Find the first match
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		name := strings.ReplaceAll(matches[1], "*", "")
		name = strings.TrimSpace(name)
		return name // The first capture group contains the description
	}
	return ""
}

func ExtractStatus(text string) string {
	// Regular expression to match the status after "# Status"
	re := regexp.MustCompile(`(?m)# Status\s+(\w+)$`)

	// Find the match
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		status := strings.TrimSpace(matches[1])
		return status
	}
	return ""
}

func ExtractDescription(input string) string {
	// Match "Description:" followed by whitespace and capture everything after it
	re := regexp.MustCompile(`Description:\s*(.*)`)
	match := re.FindStringSubmatch(input)

	if len(match) < 2 {
		return ""
	}

	// First line of description
	description := match[1]

	// Split text into lines
	lines := strings.Split(input, "\n")

	// Find the line containing Description:
	var descLineIndex int
	for i, line := range lines {
		if strings.Contains(line, "Description:") {
			descLineIndex = i
			break
		}
	}

	// Now collect all lines after the description line until we hit another key
	var additionalLines []string
	for i := descLineIndex + 1; i < len(lines); i++ {
		// If line contains a key pattern (word followed by colon), stop
		if regexp.MustCompile(`^[A-Za-z]+:`).MatchString(lines[i]) {
			break
		}
		additionalLines = append(additionalLines, lines[i])
	}

	// Combine first line with additional lines
	if len(additionalLines) > 0 {
		description += "\n" + strings.Join(additionalLines, "\n")
	}

	return strings.TrimSpace(description)
}

func WrapText(text string, width int) string {
	if width <= 0 {
		return text
	}

	var result strings.Builder
	line := strings.Builder{}
	lineLength := 0

	for _, word := range strings.SplitAfter(text, " ") {
		wordLength := len(word)

		// If adding the word exceeds the width, flush the current line
		if lineLength+wordLength > width {
			result.WriteString(line.String())
			result.WriteString("\n")
			line.Reset()
			lineLength = 0
		}

		// Add the word to the line
		line.WriteString(word)
		lineLength += wordLength
	}

	// Add any remaining content in the line buffer
	result.WriteString(line.String())
	return result.String()
}
