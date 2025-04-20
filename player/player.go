package player

import (
	"fmt"
	"strings"

	"github.com/AlexeyNilov/gorpg/textgen"
)

const (
	POVPromptStructure = `# Background:
%s
# NPC actions:
%s

Describe given above scene from the %s (%s) point of view. %s perception is %s`
)

type Player struct {
	Name        string
	Description string
	Perception  string
}

func (p *Player) GetPointOfView(tg textgen.TextGenerator, actions, background string) string {
	prompt := GetPointOfViewPrompt(*p, actions, background)
	pov, _ := tg.Generate(prompt)
	return strings.TrimSpace(pov)
}

func GetPointOfViewPrompt(player Player, actions, background string) string {
	return fmt.Sprintf(POVPromptStructure,
		background, actions, player.Name, player.Description, player.Name, player.Perception,
	)
}
