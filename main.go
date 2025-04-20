package main

import (
	"fmt"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/textgen"
)

func main() {
	scene := `Its evening, the goblin is lost in the wood. He wants to go home. But a strong wolf tries to catch him. The wolf is stalking the goblin.`
	wolf := npc.NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry"}

	textGen := &textgen.GenericTextGenerator{}

	reaction := wolf.React(textGen, scene)
	fmt.Println(reaction)

	player := player.Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
			Perception:  "poor",
		},
	}
	pov := player.GetPointOfView(textGen, reaction, scene)

	fmt.Println(pov)
}
