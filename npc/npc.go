package npc

import (
	"strings"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	LogLength        = 10
	DecisionTemplate = `# Background:
{{.Background}}
Your name is {{.NPCName}}; You are {{.NPCDescription}}

# Previous events:
{{.Events}}

# Decide what to do, be brief and realistic, focus on actions and feelings. Use 3rd point of view (use your name instead of I):`
	DescriptionUpdateTemplate = `You are the omnipotent System from a LitRPG universe, overseeing the intricately designed virtual world you've created. Update {{.NPCName}}'s description based on their actions and results. Include any significant changes to their level, HP, and skills.
# Actions:
{{.Events}}

# Initial state:
{{.NPCDescription}}

# Present your response in the following format:
Description: [Detailed Description]
`
)

type NPC struct {
	Name        string
	Description string
	Log         []string
}

func (n *NPC) LogEvent(event string) {
	n.Log = append(n.Log, event)

	// Restrict the log to the last N entries
	if len(n.Log) > LogLength {
		n.Log = n.Log[len(n.Log)-LogLength:]
	}
}

func GetPrompt(template string, npc NPC, background string) string {
	// Combine the NPC log into a single string
	events := strings.Join(npc.Log, "\n")

	data := struct {
		Background     string
		NPCName        string
		NPCDescription string
		Events         string
	}{
		Background:     background,
		NPCName:        npc.Name,
		NPCDescription: npc.Description,
		Events:         events,
	}
	return util.ParseTemplate(template, data)
}

func (n *NPC) React(tg textgen.TextGenerator, background string) string {
	prompt := GetPrompt(DecisionTemplate, *n, background)
	reaction, _ := tg.Generate(prompt)
	reaction = strings.TrimSpace(reaction)
	n.LogEvent(reaction)
	return reaction
}

func (n *NPC) UpdateDescription(tg textgen.TextGenerator, background string) {
	prompt := GetPrompt(DescriptionUpdateTemplate, *n, background)
	raw, _ := tg.Generate(prompt)
	n.Description = util.ExtractDescription(raw)
}
