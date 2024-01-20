package game

import (
	"fmt"

	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = &Game{}

type Game struct {
	systems []ecs.System
}

// WithSystems adds systems to the game.
func (g *Game) WithSystems(systems ...ecs.System) *Game {
	g.systems = append(g.systems, systems...)
	return g
}

func New() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
