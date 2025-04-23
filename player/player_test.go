package player

import (
	"strings"
	"testing"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/testutil"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/stretchr/testify/assert"
)

const mockAction = "attack"

func newTestPlayer() Player {
	mockInput := strings.NewReader(mockAction + "\n")
	return Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
		},
		Input: mockInput,
	}
}

func TestGetAction(t *testing.T) {
	player := newTestPlayer()
	action, err := player.GetAction()
	assert.Nil(t, err)
	assert.Equal(t, mockAction, action)
	assert.Equal(t, mockAction, player.Log[0])
}

func TestCreateDescription(t *testing.T) {
	player := newTestPlayer()

	assert.Equal(t, "John", player.Name)
	assert.Equal(t, "Low level goblin archer", player.Description)

	testutil.LoadEnv()
	textGen := &textgen.MockTextGenerator{Text: "Description: test desc ", Err: nil}

	player.CreateDescription(textGen, "10", "human")
	assert.Equal(t, "test desc", player.Description)
}

func TestGetName(t *testing.T) {
	mockInput := strings.NewReader("Test Name  \n")
	want := GetName(mockInput)
	assert.Equal(t, want, "Test Name")
}

func TestGetNamePanicsWhenNoName(t *testing.T) {
	// Create an input that simulates pressing "Enter" without typing anything
	mockInput := strings.NewReader("  \n")

	// Use defer and recover to catch the panic
	defer func() {
		if r := recover(); r != nil {
			// Assert that the panic message is as expected
			assert.Equal(t, "No name given", r)
		} else {
			t.Errorf("Expected a panic, but none occurred")
		}
	}()

	// Call the function, which should panic
	GetName(mockInput)
}

func TestGeneratePlayer(t *testing.T) {
	testutil.LoadEnv()
	textGen := &textgen.MockTextGenerator{Text: "Description: test desc ", Err: nil}
	player := GeneratePlayer(textGen, "John", "10", "human")
	assert.Equal(t, "John", player.Name)
	assert.Equal(t, "test desc", player.Description)
}
