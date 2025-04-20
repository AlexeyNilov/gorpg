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

func (s *Scene) Update(tg textgen.TextGenerator, reaction, pov, action string) string {
	prompt := fmt.Sprintf(`# Background: %s
# NPC actions: %s
# Player point of view: %s
# Player actions: %s
Provide brief summary, focusing only on the summary itself without any introductory or concluding phrases.`, s.Description, reaction, pov, action)
	s.Description, _ = tg.Generate(prompt)
	return s.Description
}
