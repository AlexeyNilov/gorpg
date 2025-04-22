package npc

import (
	"fmt"
	"strings"
	// "text/template"

	"github.com/AlexeyNilov/gorpg/textgen"
)

const (
	decisionPrompt  = "# Decide what to do, be brief and realistic, focus on actions and feelings. Use 3rd point of view (use your name instead of I):"
	DecisionPromptStructure = `# Background:
%s
Your name is %s; You are %s

# Previous events:
%s

%s`
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
