package scene

import (
	"log"
	"testing"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewScene(t *testing.T) {
	scene := Scene{
		Start:  "Starting position",
		System: "System prompt",
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	textGen := &textgen.MockTextGenerator{Text: "New scene", Err: nil}

	got := scene.Create(textGen)
	want := "New scene"
	assert.Equal(t, want, got)
	assert.Equal(t, want, scene.Description)
}

func TestUpdateScene(t *testing.T) {
	scene := Scene{
		Start:       "Starting position",
		System:      "System prompt",
		Description: "Old summary",
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	textGen := &textgen.MockTextGenerator{Text: "New scene", Err: nil}
	reaction := "reaction"
	action := "action"
	got := scene.Update(textGen, reaction, action)
	want := "New scene"
	assert.Equal(t, want, got)
	assert.Equal(t, want, scene.Description)
}
