package scene

import (
	"log"
	"testing"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewScene(t *testing.T) {
	scene := Scene{}

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

func TestNewNPC(t *testing.T) {
	scene := Scene{}
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	text := `Name: **Test Name**
Description: Test description
more description`
	wantName := "Test Name"
	wantDesc := `Test description
more description`
	textGen := &textgen.MockTextGenerator{Text: text, Err: nil}
	npc := scene.NewNPC(textGen, "1")
	assert.Equal(t, wantName, npc.Name)
	assert.Equal(t, wantDesc, npc.Description)
}
