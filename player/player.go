package player

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

const (
	NewPlayerTemplate = `You are the omnipotent System AKA Game Master, overseeing virtual world. Generate a brief description of a new Player named {{.Name}}. Race {{.Race}}. Randomly select their class, and include a few fitting skills appropriate to their level {{.Level}} and class. The description should include their appearance, level, and skills. Present the result in the following format:

Description: [Detailed Description, including race, class, appearance, level, HP, XP, and skills.]`
)

type Player struct {
	npc.NPC
	Input io.Reader
}

func (p *Player) CreateDescription(tg textgen.TextGenerator, level, race string) {
	data := struct {
		Level string
		Name  string
		Race  string
	}{
		Level: level,
		Name:  p.Name,
		Race:  race,
	}
	prompt := util.ParseTemplate(NewPlayerTemplate, data)
	reply, _ := tg.Generate(prompt)
	p.Description = util.ExtractDescription(reply)
}

func GetName(input io.Reader) string {
	const maxNameLength = 32 // Maximum length for the name

	fmt.Printf("Enter your name (max %d characters): ", maxNameLength)

	reader := bufio.NewReader(input)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading input:", err)
		panic(err)
	}

	// Trim whitespace and ensure the name doesn't exceed maxNameLength
	name = strings.TrimSpace(name)
	if len(name) > maxNameLength {
		name = name[:maxNameLength]
		fmt.Printf("Input truncated to %d characters.\n", maxNameLength)
	}

	if len(name) == 0 {
		panic("No name given")
	}

	return name
}

func (p *Player) GetAction() (string, error) {
	const maxActionLength = 256 // Maximum length for the action input

	// Prompt the user for input
	fmt.Printf("%s: ", p.Name)

	// Read from the injected input source
	reader := bufio.NewReader(p.Input)
	action, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Trim whitespace and ensure the action doesn't exceed maxActionLength
	action = strings.TrimSpace(action)
	if len(action) > maxActionLength {
		action = action[:maxActionLength]
		fmt.Printf("Action truncated to %d characters.\n", maxActionLength)
	}

	// Log the action
	p.LogEvent(action)
	return action, nil
}

func GeneratePlayer(tg textgen.TextGenerator, name, level, race string) Player {
	player := Player{
		NPC:   npc.NPC{Name: name},
		Input: os.Stdin,
	}
	player.CreateDescription(tg, level, race)
	return player
}
