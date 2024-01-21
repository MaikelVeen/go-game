package main

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/data"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	coordinator := ecs.NewCoordinator(
		ecs.NewEntityManager(),
		ecs.NewComponentManager(),
		ecs.NewSystemManager(),
	)

	config, err := data.LoadConfig("game.yaml")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	game := game.New(config, coordinator)

	if err := ebiten.RunGame(game); err != nil {
		slog.Error(err.Error())
	}
}
