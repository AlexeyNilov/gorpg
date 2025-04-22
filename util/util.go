package util

import (
	"bytes"
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
