package npc

import (
	"testing"

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