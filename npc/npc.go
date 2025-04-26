package npc

import (
	"fmt"
	"os"
	"strings"
	"time"

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

# Decide what to do, be very brief and realistic, focus on actions. Use 3rd point of view (use your name instead of I):`
	DescriptionUpdateTemplate = `You are the omnipotent System AKA Game Master, overseeing virtual world. Update {{.NPCName}}'s description based on their actions and results.
Include any changes to their level, HP, and skills. Repeated actions can unlock new skills or enhance existing ones.

# Actions:
{{.Events}}

# Initial state:
{{.NPCDescription}}

# Provide information using the structure below:

Description: Detailed description, including their appearance and race

# Intent
[Describe the intent]

# XP
[Current experience points/Max experience points]

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
)

type NPC struct {
	Name        string
	Description string
	Log         []string
	Status      bool
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

func (n *NPC) Die() {
	n.Name = ""
	n.Description = ""
	n.Log = nil
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
	n.Status = n.IsAlive()
}

func (n *NPC) IsAlive() bool {
	status := util.ExtractStatus(n.Description)
	if status == "Alive" {
		return true
	} else {
		return false
	}
}

func (n *NPC) AppendToFile(filename string) error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	nData := n.Name + "\n" + n.Description + "\n" + strings.Join(n.Log, "\n")
	Data := fmt.Sprintf("Timestamp: %s\n%s\n", timestamp, nData)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(Data + "\n")
	return err
}
