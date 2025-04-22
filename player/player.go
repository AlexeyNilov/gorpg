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
	NewPlayerTemplate = `You are the omnipotent System from a LitRPG universe, overseeing the intricately designed virtual world you’ve created. Generate a brief description of a new, randomly generated Player at Level {{.Level}}. Randomly select their race and class, and include a few fitting skills appropriate to their level and role. The description should include their appearance, level, and relevant abilities. Present the result in the following format:

Description: [Detailed Player Description, including race, class, appearance, level, and skills.]`
)

type Player struct {
	npc.NPC
	Input io.Reader
}

func (p *Player) Create(tg textgen.TextGenerator, level string) {
	data := struct {
		Level string
	}{
		Level: level,
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

func GeneratePlayer(tg textgen.TextGenerator, name, level string) Player {
	player := Player{
		NPC:   npc.NPC{Name: name},
		Input: os.Stdin,
	}
	player.Create(tg, level)
	return player
}
