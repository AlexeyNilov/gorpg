package storage

import (
	"os"

	"github.com/AlexeyNilov/gorpg/npc"
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
