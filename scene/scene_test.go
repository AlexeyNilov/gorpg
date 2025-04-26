package scene

import (
	"testing"

	"github.com/AlexeyNilov/gorpg/testutil"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/stretchr/testify/assert"
)

func TestNewScene(t *testing.T) {
	scene := Scene{}

	testutil.LoadEnv()
	textGen := &textgen.MockTextGenerator{Text: "New scene", Err: nil}

	got := scene.Create(textGen)
	want := "New scene"
	assert.Equal(t, want, got)
	assert.Equal(t, want, scene.Description)
}

func TestValidateAction(t *testing.T) {
	scene := Scene{
		Description: "Old summary",
	}

	testutil.LoadEnv()
	textGen := &textgen.MockTextGenerator{Text: "New scene", Err: nil}
	action := "action"
	got := scene.ValidateAction(textGen, action, "")
	want := "New scene"
	assert.Equal(t, want, got)
}

func TestNewNPC(t *testing.T) {
	scene := Scene{}
	testutil.LoadEnv()

	text := `Name: **Test Name**
Description: Test description
more description`
	textGen := &textgen.MockTextGenerator{Text: text, Err: nil}

	wantName := "Test Name"
	wantDesc := `Test description
more description`
	npc := scene.NewNPC(textGen, "1")
	assert.Equal(t, wantName, npc.Name)
	assert.Equal(t, wantDesc, npc.Description)
}

func TestGetSummary(t *testing.T) {
	scene := Scene{
		Description: "Test scene",
	}

	testutil.LoadEnv()
	textGen := &textgen.MockTextGenerator{Text: "Test action  ", Err: nil}

	got := scene.GetSummary(textGen)
	want := "Test action"
	assert.Equal(t, want, got)
}