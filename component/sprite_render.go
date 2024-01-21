package component

import "github.com/hajimehoshi/ebiten/v2"

const SpriteRenderComponentType uint8 = 1

var _ Component = (*SpriteRender)(nil)

type SpriteRender struct {
	Image *ebiten.Image
}

func (*SpriteRender) SetData(data map[string]any) error {
	return nil // TODO: Implement.
}
