package storage

import (
	"os"
	"testing"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/stretchr/testify/assert"
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

func TestLoadNPCsFromYAML(t *testing.T) {
	n := npc.NPC{Name: "Wolf", Description: "Wild wolf, powerful and hungry"}
	n.LogEvent("test")
	p := player.Player{
		NPC: npc.NPC{
			Name:        "John",
			Description: "Low level goblin archer",
		},
		Input: os.Stdin,
	}
	p.LogEvent("test\ntest")
	want := []npc.NPC{}
	want = append(want, n)
	want = append(want, p.NPC)
	_ = SaveNPCsToYAML(want, "test.yaml")
	got, _ := LoadNPCsFromYAML("test.yaml")
	assert.Equal(t, want, got)
}