package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/scene"
	"github.com/AlexeyNilov/gorpg/storage"
	"github.com/AlexeyNilov/gorpg/textgen"
)

func Loop(textGen textgen.TextGenerator, p player.Player, n npc.NPC, scene scene.Scene) {
	// Create a channel to listen for termination signals (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Infinite loop
	fmt.Print("Press Ctrl+C to exit\n\n")

	for {
		select {
		case <-stop:
			// Handle the termination signal
			fmt.Print("\nExiting gracefully...")
			return
		default:
			// Perform your loop operations here
			fmt.Print(scene.Description, "\n===================\n")

			PlayerAction, err := p.GetAction()
			NPCAction := n.React(textGen, scene.Description)

			if err != nil {
				fmt.Print("Error: ", err)
				return
			}

			fmt.Print("\n")

			scene.ValidateAction(textGen, NPCAction, PlayerAction)
			scene.UpdateBackground(textGen)
			p.LogEvent(scene.GetSummary(textGen))

			randomNumber := rand.Intn(100) + 1 // This gives a number between 1 and 100
			if randomNumber <= 50 {
				p.UpdateDescription(textGen, scene.Description)
			}

			storage.SaveState(p, n, scene)
		}
	}
}

func main() {
	textGen := &textgen.GenericTextGenerator{}

	// Define the `resume` flag
	resume := flag.Bool("resume", false, "Resume the game from the last state")

	// Parse the flags
	flag.Parse()

	// Declare variables for player, NPC, and scene
	var p player.Player
	var n npc.NPC
	var s scene.Scene

	// Use the `resume` value in your program
	if *resume {
		fmt.Println("Resuming from the last state...")
		p, n, s = storage.LoadState()
	} else {
		fmt.Println("Starting fresh...")
		s = scene.Scene{}
		s.Create(textGen)
		name := player.GetName(os.Stdin)
		p = player.GeneratePlayer(textGen, name, "1", "Human")
		n = s.NewNPC(textGen, "2")
	}

	// Main game loop
	Loop(textGen, p, n, s)
}
