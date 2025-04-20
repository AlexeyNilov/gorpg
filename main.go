package main

import (
	"fmt"
	"os"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/scene"
	"github.com/AlexeyNilov/gorpg/textgen"
)

func main() {
	player := player.Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
			Perception:  "poor",
		},
		Input: os.Stdin,
	}

	scene := scene.Scene{
		Start: fmt.Sprintf(`Its evening, the %s (%s) is lost in the wood. He wants to go home. But a strong wolf tries to catch him. The wolf is stalking.`, player.Name, player.Description),
		System: `I'm writing litRPG game system for 12 y.o. children. You are very good dungeon master. Your role is to describe scenes. Be realistic and brief. Use simple English.`,
	}
	
	wolf := npc.NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry", Perception: "sharp"}

	textGen := &textgen.GenericTextGenerator{}
	
	background := scene.Create(textGen)
	fmt.Println("Background:\n", background)

	reaction := wolf.React(textGen, background)

	pov := player.GetPointOfView(textGen, reaction, background)
	fmt.Println("POV:\n", pov)

	action := player.GetAction()
	// TODO add validation and correction
	fmt.Println(action)

	// Next cycle
	summary := scene.Update(textGen, reaction, pov, action)
	fmt.Println("Summary:\n", summary)

}
