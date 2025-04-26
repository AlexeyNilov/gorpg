package scene

import (
	"strings"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	StartScenePrompt    = `I'm writing a LitRPG novel, where the world has been dramatically transformed by the System. The protagonist is teleported to a random outdoor location. Chose between forest, desert, mountains, iceland, beach, ruins.
The setting should evoke a sense of wonder and discovery, with the protagonist alone and not in immediate danger.
Describe the location in vivid detail, using 'you' to immerse the reader as if they are experiencing the scene themselves.
Use simple English. Avoid any introductory or concluding phrases.`
	UpdateSceneTemplate = `You are a tactical AI responsible for analyzing and synthesizing environmental data in a LitRPG universe.

{{.Background}}

Analyze the provided information using the structure below:

# Location
[Describe the location, including details about the world, country, and relevant history.]

# Weather
[Detail the weather conditions.]

# Terrain
[Describe the terrain, focusing on background features and surroundings.]

# Resources
[List available resources, including food, water, and craftable materials, if present.]

# Creatures
[Identify animals or other creatures present, if any.]

# NPC
[Describe NPCs, including their state, belongings, weapons, and position relative to the terrain, if present.]

# Player
[Describe the Player, including their state, belongings, weapons, and position relative to the terrain.]

# Distances
[Specify distances between key objects, the NPC, and the Player.]

Be creative, but use clear, simple English without any introductory or concluding phrases
`
	ValidateActionTemplate = `# Background: {{.Background}}

# NPC actions: {{.NPCActions}}

# Player actions: {{.PlayerActions}}

You are the omnipotent System AKA Game Master, overseeing virtual world. Be critical and ensure the Player's actions remain grounded in their skills, stats, and level. If the Player attempts something beyond their abilities, enforce failure with humor, vividly describing the mishap. Predict and narrate the most likely outcome of the Player's actions based on their capabilities and the environment. Only describe events or NPC actions that the Player can perceive. When the Player requests information, seamlessly integrate it into your response. Avoid any introductory or concluding phrases.`
	NewNPCTemplate = `You are the omnipotent System AKA Game Master, overseeing virtual world. Generate a brief description of a new, randomly generated hostile NPC at Level {{.Level}}, tailored to fit the context of the scene: {{.Scene}}

Ensure the NPC has a clear intent, and include a few fitting skills relevant to their level and role. Present your response in the following format:

Name: [Generated NPC Name]
Description: [Detailed NPC Description, including their appearance, intent, level, HP, and skills.]`
	ActionSummaryPrompt = `Summarize the following actions in a one sentence, ensuring they are grounded in the context of the scene: {{.Scene}}`
)

type Scene struct {
	Description string
	Background  string
}

func (s *Scene) Create(tg textgen.TextGenerator) string {
	s.Description, _ = tg.Generate(StartScenePrompt)
	s.Background = s.Description
	return s.Description
}

func (s *Scene) GetSummary(tg textgen.TextGenerator) string {
	data := struct {
		Scene    string
	}{
		Scene:    s.Description,
	}
	prompt := util.ParseTemplate(ActionSummaryPrompt, data)
	summary, _ := tg.Generate(prompt)
	return strings.TrimSpace(summary)
}

func (s *Scene) UpdateBackground(tg textgen.TextGenerator) string {
	data := struct {
		Background    string
	}{
		Background:    s.Background + "\n" + s.Description,
	}
	prompt := util.ParseTemplate(UpdateSceneTemplate, data)
	s.Background, _ = tg.Generate(prompt)

	return s.Background
}

func (s *Scene) ValidateAction(tg textgen.TextGenerator, reaction, action string) string {
	data := struct {
		Background    string
		NPCActions    string
		PlayerActions string
	}{
		Background:    s.Background + "\n" + s.Description,
		NPCActions:    reaction,
		PlayerActions: action,
	}
	prompt := util.ParseTemplate(ValidateActionTemplate, data)
	s.Description, _ = tg.Generate(prompt)

	return s.Description
}

func (s *Scene) NewNPC(tg textgen.TextGenerator, level string) npc.NPC {
	data := struct {
		Level string
		Scene string
	}{
		Level: level,
		Scene: s.Description,
	}
	prompt := util.ParseTemplate(NewNPCTemplate, data)
	reply, _ := tg.Generate(prompt)
	return npc.NPC{Name: util.ExtractName(reply), Description: util.ExtractDescription(reply)}
}
