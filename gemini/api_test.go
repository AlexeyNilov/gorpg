package gemini

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGenerate(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	intro := `I will give you short texts describing situations from 1st point of view. You role to give very brief summary focused on actions.`
	action := `As I near the stream, I make my decision. I leap down the slope, arms flailing to keep balance. My feet slip on the slick earth, but I crash forward into the freezing water, the shock stealing my breath. I scramble to my feet, ignoring the cold biting into my skin, and start running downstream. The current’s pull is strong, but it’s better than teeth.`
	prompt := intro + "/n" + action
	got := GenerateText(prompt)
	want := "text"

	if got != want {
		t.Errorf("Want %v, got %v", want, got)
	}
}
