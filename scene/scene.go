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
	prompt := s.System + "\n" + s.Start + "\nImagine and describe background"
	s.Description, _ = tg.Generate(prompt)
	return s.Description
}

func (s *Scene) Update(tg textgen.TextGenerator, reaction, pov, action string) string {
	prompt := fmt.Sprintf(`# Background: %s
# NPC actions: %s
# Player point of view: %s
# Player actions: %s
Provide brief summary`, s.Start, reaction, pov, action)
	s.Start, _ = tg.Generate(prompt)
	return s.Start
}
