package player

import (
	"log"
	"testing"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func newTestPlayer() Player {
	return Player{Name: "John", Description: "Low level goblin archer", Perception: "poor"}
}

func TestGetPointOfViewPrompt(t *testing.T) {
	testPlayer := newTestPlayer()
	background := "Test scene description"
	npcActions := "Test action"
	got := GetPointOfViewPrompt(testPlayer, npcActions, background)
	want := `# Background:
Test scene description
# NPC actions:
Test action

Describe given above scene from the John (Low level goblin archer) point of view. John perception is poor`
	assert.Equal(t, want, got)
}

func TestGetPOV(t *testing.T) {
	testPlayer := newTestPlayer()
	background := "Test scene description"
	npcActions := "Test action"

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	textGen := &textgen.MockTextGenerator{Text: "test pov    ", Err: nil}
	got := testPlayer.GetPointOfView(textGen, npcActions, background)
	want := "test pov"
	assert.Equal(t, got, want)
}
