package data

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AssetConfig struct {
	SpriteSheets []SpriteSheetConfig `yaml:"spriteSheets"`
}

type SpriteSheetConfig struct {
	Name      string `yaml:"name"`
	Path      string `yaml:"path"`
	Cols      int    `yaml:"cols"`
	Rows      int    `yaml:"rows"`
	FrameSize int    `yaml:"frameSize"`
}

func LoadAssetConfig(filename string) (*AssetConfig, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config AssetConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
