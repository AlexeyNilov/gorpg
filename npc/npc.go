package npc

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/AlexeyNilov/gorpg/textgen"
)

const (
	LogLength               = 10
	DecisionPromptStructure = `# Background:
{{.Background}}
Your name is {{.NPCName}}; You are {{.NPCDescription}}

# Previous events:
{{.Events}}

# Decide what to do, be brief and realistic, focus on actions and feelings. Use 3rd point of view (use your name instead of I):`
)

type Prompt struct {
	Background     string
	NPCName        string
	NPCDescription string
	Events         string
}

type NPC struct {
	Name        string
	Description string
	Log         []string
	Perception  string
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

	// Define the prompt structure
	prompt := Prompt{
		Background:     background,
		NPCName:        npc.Name,
		NPCDescription: npc.Description,
		Events:         events,
	}

	// Parse the template
	tpl, err := template.New("Decision").Parse(DecisionPromptStructure)
	if err != nil {
		panic(err)
	}

	// Use a bytes.Buffer to capture the output
	var buf bytes.Buffer
	err = tpl.Execute(&buf, prompt)
	if err != nil {
		panic(err)
	}

	// Return the captured output as a string
	return buf.String()
}

func (n *NPC) React(tg textgen.TextGenerator, background string) string {
	prompt := GetDecisionPrompt(*n, background)
	reaction, _ := tg.Generate(prompt)
	n.LogEvent(reaction)
	return strings.TrimSpace(reaction)
}
