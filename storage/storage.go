package storage

import (
	"os"

	"github.com/AlexeyNilov/gorpg/npc"
	"github.com/AlexeyNilov/gorpg/player"
	"github.com/AlexeyNilov/gorpg/scene"
	"gopkg.in/yaml.v3"
)

func SaveNPCsToYAML(npcs []npc.NPC, filename string) error {
	data, err := yaml.Marshal(npcs)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644) // Replaces ioutil.WriteFile
}

func LoadNPCsFromYAML(filename string) ([]npc.NPC, error) {
	data, err := os.ReadFile(filename) // Replaces ioutil.ReadFile
	if err != nil {
		return nil, err
	}
	var npcs []npc.NPC
	err = yaml.Unmarshal(data, &npcs)
	return npcs, err
}

func SaveSceneToYAML(s scene.Scene, filename string) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadSceneFromYAML(filename string) (scene.Scene, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return scene.Scene{}, err
	}
	var scene scene.Scene
	err = yaml.Unmarshal(data, &scene)
	return scene, err
}

func SaveState(p player.Player, n npc.NPC, s scene.Scene) {
	npcs := []npc.NPC{n, p.NPC}
	_ = SaveNPCsToYAML(npcs, "data/npc.yaml")
	_ = SaveSceneToYAML(s, "data/scene.yaml")
	_ = s.AppendToFile("data/game.log")
}

func LoadState() (player.Player, npc.NPC, scene.Scene ){
	npcs, _ := LoadNPCsFromYAML("data/npc.yaml")
	n := npcs[0]
	p := player.Player{
		NPC: npcs[1],
		Input: os.Stdin,
	}
	scene, _ := LoadSceneFromYAML("data/scene.yaml")

	return p, n, scene 

}
