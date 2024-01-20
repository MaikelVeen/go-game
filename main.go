package main

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game.New()); err != nil {
		slog.Error(err.Error())
	}
}
