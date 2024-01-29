package spriterenderer

import (
	"fmt"

	"github.com/MaikelVeen/go-game/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Type uint = 1
	Slug      = "spriteRenderer"
)

type SpriteRenderer struct {
	SpriteSheet assets.SpriteSheet
	// Dictates which sprite to render by default.
	X, Y int
}

// SetData implements Component.
func (s *SpriteRenderer) SetData(data map[string]any) error {
	spriteSheetName, exists := data["spriteSheet"].(string)
	if !exists {
		return fmt.Errorf("could not set sprite render data: missing name")
	}

	spriteSheet, exists := assets.GlobalAssetRegistry.SpriteSheets[spriteSheetName]
	if !exists {
		return fmt.Errorf("could not set sprite render data: sprite sheet %s not found", spriteSheetName)
	}
	s.SpriteSheet = spriteSheet

	x, exists := data["x"].(int)
	if !exists {
		return fmt.Errorf("could not set sprite render data: missing x")
	}
	s.X = x

	y, exists := data["y"].(int)
	if !exists {
		return fmt.Errorf("could not set sprite render data: missing y")
	}
	s.Y = y

	return nil
}

func (s *SpriteRenderer) GetSprite() *ebiten.Image {
	return s.SpriteSheet.Sprite(s.X, s.Y)
}
