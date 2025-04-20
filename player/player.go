package player

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/AlexeyNilov/gorpg/npc"
)

type Player struct {
	npc.NPC
	Input io.Reader
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
