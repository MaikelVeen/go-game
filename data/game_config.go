package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GameConfig struct {
	Entities []EntityConfig `yaml:"entities"`
}

type EntityConfig struct {
	Name       string            `yaml:"name"`
	Components []ComponentConfig `yaml:"components"`
}

type ComponentConfig struct {
	Type string         `yaml:"type"`
	Data map[string]any `yaml:"data"`
}

func LoadGameConfig(filename string) (*GameConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config GameConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
