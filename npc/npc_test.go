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

func TestDescribe(t *testing.T) {
	npc := newTestNPC()
	got := npc.Describe()
	want := "Wild wolf, powerful and hungry"
	assert.Equal(t, want, got)
}

func TestLogEvent(t *testing.T) {
	npc := newTestNPC()
	npc.LogEvent("Woke up")
	want := []string{"Woke up"}
	assert.Equal(t, want, npc.Log)
}

func TestGetPrompt(t *testing.T) {
	background := "Dense forest, night"

	npc := newTestNPC()
	npc.LogEvent("Woke up")
	npc.LogEvent("Sniff air")

	got := GetPrompt(npc, background)

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
