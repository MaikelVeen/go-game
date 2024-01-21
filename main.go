package main

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/assets"
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

	gameConfig, err := data.LoadGameConfig("game.yaml")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	assetConfig, err := data.LoadAssetConfig("assets.yaml")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	assets.GlobalAssetRegistry = assets.NewAssetRegistry()
	if err := assets.GlobalAssetRegistry.LoadAssets(assets.StaticSpritesFS, assetConfig); err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	game := game.New(gameConfig, coordinator)

	if err := ebiten.RunGame(game); err != nil {
		slog.Error(err.Error())
	}
}
