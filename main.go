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

func Loop(textGen textgen.TextGenerator, scene scene.Scene, n npc.NPC, p player.Player) {
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

			NPCAction := n.React(textGen, scene.Description)

			PlayerAction, err := p.GetAction()
			if err != nil {
				fmt.Print("Error: ", err)
				return
			}

			fmt.Print("\n")

			scene.Update(textGen, NPCAction, PlayerAction)
			p.LogEvent(scene.GetSummary(textGen))

			randomNumber := rand.Intn(100) + 1 // This gives a number between 1 and 100
			if randomNumber <= 20 {
				p.UpdateDescription(textGen, scene.Description)
			}
			_ = npc.AppendToFile("log/player.log", p.NPC)
			_ = npc.AppendToFile("log/npc.log", n)
		}
	}
}

func main() {
	textGen := &textgen.GenericTextGenerator{}

	scene := scene.Scene{}
	scene.Create(textGen)
	name := player.GetName(os.Stdin)
	player := player.GeneratePlayer(textGen, name, "1", "Human")
	npc := scene.NewNPC(textGen, "2")

	Loop(textGen, scene, npc, player)

}
