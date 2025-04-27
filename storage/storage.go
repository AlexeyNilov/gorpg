package storage

import (
	"os"

	"gopkg.in/yaml.v3"
)

// SaveToYAML saves data to a YAML file.
func SaveToYAML(data any, filename string) error {
	out, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, out, 0644)
}

// LoadFromYAML loads data from a YAML file into the provided output parameter.
func LoadFromYAML(filename string, out any) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, out)
}
