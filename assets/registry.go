package assets

import (
	"bytes"
	"embed"
	"fmt"
	"image"
	"image/png"

	"github.com/MaikelVeen/go-game/data"
	"github.com/hajimehoshi/ebiten/v2"
)

var GlobalAssetRegistry *AssetRegistry

type AssetRegistry struct {
	SpriteSheets map[string]SpriteSheet
}

func NewAssetRegistry() *AssetRegistry {
	return &AssetRegistry{
		SpriteSheets: make(map[string]SpriteSheet),
	}
}

// LoadAssets loads all assets from the given embed.FS.
func (ar *AssetRegistry) LoadAssets(fs embed.FS, config *data.AssetConfig) error {
	// Load sprite sheets.
	for _, cfg := range config.SpriteSheets {
		// Load image.
		img, err := loadFile(fs, cfg.Path)
		if err != nil {
			return fmt.Errorf("failed to load sprite %s: %w", cfg.Path, err)
		}

		// Create ebiten image.
		ebitenImage := ebiten.NewImageFromImage(img)

		// Create sprite sheet.
		spriteSheet := NewSpriteSheet(ebitenImage,
			SliceImage(cfg.Cols, cfg.Rows, cfg.FrameSize, cfg.FrameSize))

		// Add sprite sheet to registry.
		ar.SpriteSheets[cfg.Name] = *spriteSheet
	}

	return nil
}

func loadFile(fs embed.FS, name string) (image.Image, error) {
	content, err := StaticSpritesFS.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to load png: %w", err)
	}

	img, err := png.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to decode png: %w", err)
	}

	return img, nil
}
