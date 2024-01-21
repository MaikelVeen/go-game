package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	UpKeys    = KeySlice{ebiten.KeyW, ebiten.KeyUp}
	DownKeys  = KeySlice{ebiten.KeyS, ebiten.KeyDown}
	LeftKeys  = KeySlice{ebiten.KeyA, ebiten.KeyLeft}
	RightKeys = KeySlice{ebiten.KeyD, ebiten.KeyRight}
)

type KeySlice []ebiten.Key

// PressedAny returns true if any of the keys in the slice are pressed.
func (s KeySlice) PressedAny() bool {
	for _, key := range s {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}

	return false
}
