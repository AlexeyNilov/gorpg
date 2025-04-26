package scene

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	StartScenePrompt = `I'm writing a LitRPG novel, where the world has been dramatically transformed by the System. The protagonist is teleported to a random location. Chose one of: forest, desert, mountains, iceland, beach, city ruins, swamp, canyon, lake.
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

# Time
[Time of day and celestial bodies]

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

# Description: {{.Description}}

# Actions: {{.Actions}}

You are the omnipotent System AKA Game Master, overseeing virtual world. Be helpful and ensure the Player/NPC gets what he/she wants.
Predict and narrate the most likely outcome of the Player/NPC actions based on their capabilities and the environment.
When the Player requests information, seamlessly integrate it into your response. Avoid any introductory or concluding phrases.`
	NewNPCTemplate = `You are the omnipotent System AKA Game Master, overseeing virtual world. Generate a brief description of a new hostile|friendly|chaotic|neutral NPC at Level {{.Level}}, tailored to fit the context of the scene: {{.Scene}}

Ensure the NPC has a clear intent, and include a few fitting skills relevant to their level and role.
Provide information using the structure below:

Name: [Generated NPC Name]
Description: Detailed NPC Description, including their appearance and race

# Intent
[Describe the intent]

# Level
[Current level]

# HP
[Current health points/Max health points]

# Status
[Dead|Alive]

# Skills
[List skills]

# Inventory
[List items and weapons]
`
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
		Scene string
	}{
		Scene: s.Description,
	}
	prompt := util.ParseTemplate(ActionSummaryPrompt, data)
	summary, _ := tg.Generate(prompt)
	return strings.TrimSpace(summary)
}

func (s *Scene) UpdateBackground(tg textgen.TextGenerator) string {
	data := struct {
		Background string
	}{
		Background: s.Background + "\n" + s.Description,
	}
	prompt := util.ParseTemplate(UpdateSceneTemplate, data)
	s.Background, _ = tg.Generate(prompt)

	return s.Background
}

func (s *Scene) ValidateAction(tg textgen.TextGenerator, Action, Description string) string {
	data := struct {
		Background  string
		Actions     string
		Description string
	}{
		Background:  s.Background + "\n" + s.Description,
		Actions:     Action,
		Description: Description,
	}
	prompt := util.ParseTemplate(ValidateActionTemplate, data)
	validation, _ := tg.Generate(prompt)

	return validation
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
	n := npc.NPC{Name: util.ExtractName(reply), Description: util.ExtractDescription(reply)}
	n.Status = n.IsAlive()
	return n
}

// Function to append Scene data with a timestamp to a file
func (scene *Scene) AppendToFile(filename string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	Data := fmt.Sprintf("Timestamp: %s\n%s\n", timestamp, scene.Description)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(Data + "\n")
	return err
}
