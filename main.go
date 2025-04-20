package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/scene"
	"github.com/AlexeyNilov/gorpg/textgen"
)

func Loop(textGen textgen.TextGenerator, scene scene.Scene, npc npc.NPC, player player.Player) {
	// Create a channel to listen for termination signals (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Infinite loop
	fmt.Println("Press Ctrl+C to exit.")
	for {
		select {
		case <-stop:
			// Handle the termination signal
			fmt.Print("\nExiting gracefully...")
			return
		default:
			// Perform your loop operations here
			fmt.Print("Background:\n", scene.Description, "\n===================")

			reaction := npc.React(textGen, scene.Description)
			fmt.Print("NPC Reaction:\n", reaction, "\n===================")

			pov := player.GetPointOfView(textGen, reaction, scene.Description)
			fmt.Print("Player POV:\n", pov, "\n===================")

			action := player.GetAction()
			// TODO add validation and correction

			fmt.Print("\n\n\n")

			// Next cycle
			scene.Update(textGen, reaction, pov, action)
		}
	}
}

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
		Start:  fmt.Sprintf(`Its evening, the %s (%s) is lost in the wood. He wants to go home. But a strong wolf tries to catch him. The wolf is stalking.`, player.Name, player.Description),
		System: `I'm writing litRPG game system for 12 y.o. children. You are very good dungeon master. Your role is to describe scenes. Be realistic and brief. Use simple English.`,
	}

	wolf := npc.NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry", Perception: "sharp"}

	textGen := &textgen.GenericTextGenerator{}

	scene.Create(textGen)

	Loop(textGen, scene, wolf, player)

}
