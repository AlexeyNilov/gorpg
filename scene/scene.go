package scene

import (
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	StartScenePrompt    = `I'm writing a LitRPG novel set in the System Apocalypse universe, where the world has been dramatically transformed by the System. The protagonist is teleported to a random outdoor location. The setting should evoke a sense of wonder and discovery, with the protagonist alone and not in immediate danger. Describe the location in vivid detail, using 'you' to immerse the reader as if they are experiencing the scene themselves. Use simple English.`
	UpdateSceneTemplate = `# Background: {{.Background}}
# NPC actions: {{.NPCActions}}
# Player actions: {{.PlayerActions}}

You are the omnipotent System from a LitRPG universe, overseeing a virtual world of your creation. Be critical and ensure the Player's actions remain grounded in their skills, stats, and level. If the Player attempts something beyond their abilities, enforce failure with humor, vividly describing the mishap. Predict and narrate the most likely outcome of the Player's actions based on their capabilities and the environment. Only describe events or NPC actions that the Player can perceive. When the Player requests information, seamlessly integrate it into your response. Avoid any introductory or concluding phrases.`
)

type Scene struct {
	Description string
}

func (s *Scene) Create(tg textgen.TextGenerator) string {
	s.Description, _ = tg.Generate(StartScenePrompt)
	return s.Description
}

func (s *Scene) Update(tg textgen.TextGenerator, reaction, action string) string {
	data := struct {
		Background    string
		NPCActions    string
		PlayerActions string
	}{
		Background:    s.Description,
		NPCActions:    reaction,
		PlayerActions: action,
	}
	prompt := util.ParseTemplate(UpdateSceneTemplate, data)
	s.Description, _ = tg.Generate(prompt)
	return s.Description
}
