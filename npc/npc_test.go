package npc

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/testutil"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/stretchr/testify/assert"
)

const mockText = `# Background:
Dense forest, night
Your name is Wolf; You are Wild wolf, powerful and hungry

# Previous events:
Woke up
Sniff air

# Decide what to do, be brief and realistic, focus on actions and feelings. Use 3rd point of view (use your name instead of I):`

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

func TestLogEventLength(t *testing.T) {
	npc := newTestNPC()
	for range 20 {
		npc.LogEvent("Next event")
	}
	want := LogLength
	assert.Equal(t, want, len(npc.Log))
}

func TestUpdateDescription(t *testing.T) {
	background := "Dense forest, night"
	npc := newTestNPC()

	testutil.LoadEnv()

	textGen := &textgen.MockTextGenerator{Text: "Description: Do something ", Err: nil}
	npc.UpdateDescription(textGen, background)

	want := "Do something"
	assert.Equal(t, want, npc.Description)
}

func TestGetPrompt(t *testing.T) {
	background := "Dense forest, night"

	npc := newTestNPC()
	npc.LogEvent("Woke up")
	npc.LogEvent("Sniff air")

	got := GetPrompt(DecisionTemplate, npc, background)

	want := mockText
	assert.Equal(t, want, got)
}

func TestReact(t *testing.T) {
	background := "Dense forest, night"
	npc := newTestNPC()

	testutil.LoadEnv()

	textGen := &textgen.MockTextGenerator{Text: "Do something    ", Err: nil}
	got := npc.React(textGen, background)

	want := "Do something"
	assert.Equal(t, want, got)
	assert.Equal(t, want, npc.Log[0])
}
