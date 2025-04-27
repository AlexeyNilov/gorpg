package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveToYAML(t *testing.T) {
	n := struct {
		Name        string
		Description string
	}{Name: "Test", Description: "Test description"}

	err := SaveToYAML(n, "test.yaml")
	assert.NoError(t, err)
}

func TestLoadFromYAML(t *testing.T) {
	want := struct {
		Name        string
		Description string
	}{Name: "Test", Description: "Test description"}

	err := SaveToYAML(want, "test.yaml")
	assert.NoError(t, err)

	var got struct {
		Name        string
		Description string
	}
	err = LoadFromYAML("test.yaml", &got)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
