package storage

import (
	"os"
	"testing"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
)

func TestSaveNPCsToYAML(t *testing.T) {
	n := npc.NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry"}
	p := player.Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
		},
		Input: os.Stdin,
	}
	npcs := []npc.NPC{}
	npcs = append(npcs, n)
	npcs = append(npcs, p.NPC)
	_ = SaveNPCsToYAML(npcs, "test.yaml")
}
