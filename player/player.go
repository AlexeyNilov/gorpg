package player

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	NewPlayerTemplate = `You are the omnipotent System from a LitRPG universe, overseeing the intricately designed virtual world you’ve created. Generate a brief description of a new Player named {{.Name}}. Race {{.Race}}. Randomly select their class, and include a few fitting skills appropriate to their level {{.Level}} and role. The description should include their appearance, level, and relevant abilities. Present the result in the following format:

Description: [Detailed Player Description, including race, class, appearance, level, and skills.]`
)

type Player struct {
	npc.NPC
	Input io.Reader
}

func (p *Player) Create(tg textgen.TextGenerator, level, race string) {
	data := struct {
		Level string
		Name string
		Race string
	}{
		Level: level,
		Name: p.Name,
		Race: race,
	}
	prompt := util.ParseTemplate(NewPlayerTemplate, data)
	reply, _ := tg.Generate(prompt)
	p.Description = util.ExtractDescription(reply)
}

func GetName() string {
	fmt.Print("Enter your name: ")

	// Read from the injected input source
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(name)
}

func (p *Player) GetAction() string {
	// Prompt the user for input
	fmt.Printf("%s, enter your action: ", p.Name)

	// Read from the injected input source
	reader := bufio.NewReader(p.Input)
	action, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading input:", err)
		return ""
	}

	action = strings.TrimSpace(action)
	p.LogEvent(action)
	return action
}

func GeneratePlayer(tg textgen.TextGenerator, name, level, race string) Player {
	player := Player{
		NPC:   npc.NPC{Name: name},
		Input: os.Stdin,
	}
	player.Create(tg, level, race)
	return player
}

// Function to append NPC data with a timestamp to a file
func AppendToFile(name string, p Player) error {
	// Get the current timestamp in a human-readable format
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Format the NPC data into a human-readable string with the timestamp
	npcData := fmt.Sprintf("Timestamp: %s\nName: %s\nDescription: %s\nLog:\n", timestamp, p.Name, p.Description)
	for _, event := range p.Log {
		npcData += fmt.Sprintf("  - %s\n", event)
	}

	// Open the file npc.log in append mode, create it if it doesn't exist
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the formatted NPC data to the file
	_, err = file.WriteString(npcData + "\n")
	return err
}