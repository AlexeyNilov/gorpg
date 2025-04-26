package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/scene"
	"github.com/AlexeyNilov/gorpg/storage"
	"github.com/AlexeyNilov/gorpg/textgen"
	"github.com/AlexeyNilov/gorpg/util"
)

func Loop(textGen textgen.TextGenerator, p player.Player, n npc.NPC, scene scene.Scene) {
	// Create a channel to listen for termination signals (Ctrl+C)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Infinite loop
	fmt.Print("Press Ctrl+C to exit\n\n")

	var NPCAction string
	// var Summary string
	var NPCValidation string
	var PlayerValidation string

	for {
		select {
		case <-stop:
			// Handle the termination signal
			fmt.Print("\nExiting gracefully...")
			return
		default:
			// Perform your loop operations here
			wrappedText := util.WrapText(scene.Description, 96)
			// wrappedText = scene.Description
			fmt.Print(wrappedText, "\n===================\n")

			PlayerAction, err := p.GetAction()
			if PlayerAction == "" {
				fmt.Print("\nNo action, exiting...")
				return
			}

			if n.Status {
				NPCAction = n.React(textGen, scene.Description)
			} else {
				fmt.Print("\n", n.Name, " is dead\n")
				n.Die()
				NPCAction = ""
			}

			if err != nil {
				fmt.Print("Error: ", err)
				return
			}

			fmt.Print("\n")

			NPCValidation = scene.ValidateAction(textGen, NPCAction, n.Description)
			PlayerValidation = scene.ValidateAction(textGen, PlayerAction, p.Description)
			scene.Description = PlayerValidation
			scene.UpdateBackground(textGen)
			// Summary = scene.GetSummary(textGen)
			p.LogEvent(PlayerValidation)
			n.LogEvent(NPCValidation)

			if n.Status {
				n.UpdateDescription(textGen, NPCValidation)
			}

			p.UpdateDescription(textGen, PlayerValidation)
			storage.SaveState(p, n, scene)

			if !p.Status {
				fmt.Print("\nYou are dead. Game over.")
				return
			}
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
		n = s.NewNPC(textGen, "1")
		storage.SaveState(p, n, s)
	}

	// Main game loop
	Loop(textGen, p, n, s)
}
