package scene

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/textgen"
)

type Scene struct {
	Start       string
	System      string
	Description string
}

func (s *Scene) Create(tg textgen.TextGenerator) string {
	prompt := s.System + "\n" + s.Start + `\nDescribe background of the given scene, focusing only on the description itself without any introductory or concluding phrases.`
	s.Description, _ = tg.Generate(prompt)
	return s.Description
}

func (s *Scene) Update(tg textgen.TextGenerator, reaction, action string) string {
	prompt := fmt.Sprintf(`# Background: %s
# NPC actions: %s
# Player actions: %s

You are almighty System (from LitRPG books) that created this virtual world. Be critical.
Make sure Player actions are realistic.
If player tries to do something impossible for his skill and level - fail it.
Use humor describing the failure.
Predict and describe most probable outcome of the actions.
Do not describe NPC actions that are not visible to the Player.
If Player requests some information include it into the description.
Do not use any introductory or concluding phrases.`, s.Description, reaction, action)
	s.Description, _ = tg.Generate(prompt)
	return s.Description
}
