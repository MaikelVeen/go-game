package components

import "github.com/hajimehoshi/ebiten/v2"

const SpriteRenderComponentType uint8 = 1

type SpriteRender struct {
	Image *ebiten.Image
}
