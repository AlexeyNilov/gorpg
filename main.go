package main

import (
	"fmt"
	"math/rand"
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
	fmt.Print("Press Ctrl+C to exit\n")
	fmt.Print("Story begins\n", "===================\n")

	for {
		select {
		case <-stop:
			// Handle the termination signal
			fmt.Print("\nExiting gracefully...")
			return
		default:
			// Perform your loop operations here
			fmt.Print(scene.Description, "\n===================\n")

			reaction := npc.React(textGen, scene.Description)

			action := player.GetAction()

			fmt.Print("\n")

			// Next cycle
			scene.Update(textGen, reaction, action)

			randomNumber := rand.Intn(100) + 1 // This gives a number between 1 and 100
			if randomNumber <= 10 {
				player.UpdateDescription(textGen, scene.Description)
				fmt.Print("\nPlayer updated\n", player.Description, "\n\n")
			}
		}
	}
}

func main() {
	textGen := &textgen.GenericTextGenerator{}

	scene := scene.Scene{}
	scene.Create(textGen)
	name := player.GetName()
	player := player.GeneratePlayer(textGen, name, "1")
	npc := scene.NewNPC(textGen, "2")

	Loop(textGen, scene, npc, player)

}
