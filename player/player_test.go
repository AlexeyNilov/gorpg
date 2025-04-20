package player

import (
	"strings"
	"testing"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/stretchr/testify/assert"
)

func newTestPlayer() Player {
	mockInput := strings.NewReader("attack\n")
	return Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
			Perception:  "poor",
		},
		Input: mockInput,
	}
}

func TestGetAction(t *testing.T) {
	// Arrange: Set up a Player
	player := newTestPlayer()

	action := player.GetAction()

	// Assert: Verify the result
	assert.Equal(t, "attack", action)
}
