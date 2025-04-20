package npc

import (
	"log"
	"testing"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a sample NPC for testing
func newTestNPC() NPC {
	return NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry"}
}

func TestLogEvent(t *testing.T) {
	npc := newTestNPC()
	npc.LogEvent("Woke up")
	want := []string{"Woke up"}
	assert.Equal(t, want, npc.Log)
}

func TestGetDecisionPrompt(t *testing.T) {
	background := "Dense forest, night"

	npc := newTestNPC()
	npc.LogEvent("Woke up")
	npc.LogEvent("Sniff air")

	got := GetDecisionPrompt(npc, background)

	want := `# Background:
Dense forest, night
Your name is Wolf; You are Wild wolf, powerful and hungry

# Previous events:
Woke up
Sniff air

# Decide what to do, be brief and realistic, focus on actions and feelings:`

	assert.Equal(t, want, got)
}

func TestReact(t *testing.T) {
	background := "Dense forest, night"
	npc := newTestNPC()

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	textGen := &textgen.MockTextGenerator{Text: "Do something    ", Err: nil}
	// textGen := &textgen.GenericTextGenerator{}
	got := npc.React(textGen, background)

	want := "Do something"
	assert.Equal(t, want, got)
}


func newTestPlayer() NPC {
	return NPC{Name: "John", Description: "Low level goblin archer", Perception: "poor"}
}

func TestGetPointOfViewPrompt(t *testing.T) {
	testPlayer := newTestPlayer()
	background := "Test scene description"
	npcActions := "Test action"
	got := GetPointOfViewPrompt(testPlayer, npcActions, background)
	want := `# Background:
Test scene description
# Actions:
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