package npc

import (
	"strings"

	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	LogLength              = 10
	DecisionPromptTemplate = `# Background:
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
	// Append the new event to the log
	n.Log = append(n.Log, event)

	// Restrict the log to the last 10 entries
	if len(n.Log) > LogLength {
		n.Log = n.Log[len(n.Log)-LogLength:]
	}
}

func GetDecisionPrompt(npc NPC, background string) string {
	// Combine the NPC log into a single string
	events := strings.Join(npc.Log, "\n")

	prompt := struct {
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
	return util.ParseTemplate(DecisionPromptTemplate, prompt)
}

func (n *NPC) React(tg textgen.TextGenerator, background string) string {
	prompt := GetDecisionPrompt(*n, background)
	reaction, _ := tg.Generate(prompt)
	n.LogEvent(reaction)
	return strings.TrimSpace(reaction)
}

func (n *NPC) UpdateDescription(tg textgen.TextGenerator, background string) {
	events := strings.Join(n.Log, "\n")

	data := struct {
		Background     string
		NPCName        string
		NPCDescription string
		Events         string
	}{
		Background:     background,
		NPCName:        n.Name,
		NPCDescription: n.Description,
		Events:         events,
	}
	prompt := util.ParseTemplate(DescriptionUpdateTemplate, data)
	raw, _ := tg.Generate(prompt)
	n.Description = util.ExtractDescription(raw)
}
