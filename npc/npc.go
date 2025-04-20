package npc

import (
	"fmt"
	"strings"

	"github.com/AlexeyNilov/gorpg/textgen"
)

const (
	decisionPrompt  = "# Decide what to do, be brief and realistic, focus on actions and feelings:"
	DecisionPromptStructure = `# Background:
%s
Your name is %s; You are %s

# Previous events:
%s

%s`
	POVPromptStructure = `# Background:
%s
# Actions:
%s

Describe given above scene from the %s (%s) point of view. %s perception is %s`
)

type NPC struct {
	Name        string
	Description string
	Log         []string
	Perception  string
}

func (n *NPC) LogEvent(event string) {
	n.Log = append(n.Log, event)
}

func GetDecisionPrompt(npc NPC, background string) string {
	events := strings.Join(npc.Log, "\n")
	return fmt.Sprintf(DecisionPromptStructure,
		background, npc.Name, npc.Description, events, decisionPrompt,
	)
}

func (n *NPC) React(tg textgen.TextGenerator, background string) string {
	prompt := GetDecisionPrompt(*n, background)
	reaction, _ := tg.Generate(prompt)
	n.LogEvent(reaction)
	return strings.TrimSpace(reaction)
}

func (n *NPC) GetPointOfView(tg textgen.TextGenerator, actions, background string) string {
	prompt := GetPointOfViewPrompt(*n, actions, background)
	pov, _ := tg.Generate(prompt)
	return strings.TrimSpace(pov)
}

func GetPointOfViewPrompt(npc NPC, actions, background string) string {
	return fmt.Sprintf(POVPromptStructure,
		background, actions, npc.Name, npc.Description, npc.Name, npc.Perception,
	)
}
