package npc

import (
	"fmt"
	"strings"

	"github.com/AlexeyNilov/gorpg/textgen"
)

const (
	decisionPrompt  = "# Decide what to do, be brief and realistic, focus on actions and feelings:"
	promptStructure = `# Background:
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
}

func (n *NPC) Describe() string {
	return n.Description
}

func (n *NPC) LogEvent(event string) {
	n.Log = append(n.Log, event)
}

func GetPrompt(npc NPC, background string) string {
	events := strings.Join(npc.Log, "\n")
	return fmt.Sprintf(promptStructure,
		background, npc.Name, npc.Describe(), events, decisionPrompt,
	)
}

func (n *NPC) React(tg textgen.TextGenerator, background string) string {
	prompt := GetPrompt(*n, background)
	reaction, _ := tg.Generate(prompt)
	return strings.TrimSpace(reaction)
}
