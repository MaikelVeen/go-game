package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/MaikelVeen/go-game/assets"
	"github.com/MaikelVeen/go-game/data"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/game"
	"github.com/MaikelVeen/go-game/timer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lmittmann/tint"
)

const (
	screenWidth  = 1920
	screenHeight = 1080

	configPath       = "game.yaml"
	assetsConfigPath = "assets.yaml"
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	logger := slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.TimeOnly,
		}),
	)
	slog.SetDefault(logger)

	coordinator := ecs.NewCoordinator(
		ecs.NewEntityManager(),
		ecs.NewComponentManager(),
		ecs.NewSystemManager(),
	)

	gameConfig, err := data.LoadGameConfig(configPath)
	if err != nil {
		slog.Error("Could not load game config", "error", err)
		return
	}

	if err := loadAssets(assetsConfigPath); err != nil {
		slog.Error("Could not load assets", "error", err)
		return
	}

	game, err := game.New(gameConfig, coordinator)
	if err != nil {
		slog.Error("Could not create game", "error", err)
		return
	}

	// Init game.
	if err := game.Init(); err != nil {
		slog.Error("Could not init game", "error", err)
		return
	}

	if err := ebiten.RunGame(game); err != nil {
		slog.Error(err.Error())
	}
}

// loadAssets creates a new AssetRegistry and loads all assets into it.
func loadAssets(cfgPath string) error {
	timer.Timer("loadAssets")()

	assetConfig, err := data.LoadAssetConfig(cfgPath)
	if err != nil {
		return err
	}

	assets.GlobalAssetRegistry = assets.NewAssetRegistry()
	if err := assets.GlobalAssetRegistry.LoadAssets(assets.StaticSpritesFS, assetConfig); err != nil {
		return err
	}

	return nil
}
