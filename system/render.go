package system

import (
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

// Ensure RenderSystem implements System.
var _ ecs.System = &RenderSystem{}

type RenderSystem struct {
	// Reference to screen.
	screen *ebiten.Image
}

func NewRenderSystem(screen *ebiten.Image) *RenderSystem {
	return &RenderSystem{
		screen: screen,
	}
}

func (s *RenderSystem) Update() {

}
