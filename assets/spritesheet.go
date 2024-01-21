package assets

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	Image *ebiten.Image
	Slice [][]image.Rectangle
}

func NewSpriteSheet(image *ebiten.Image, slice [][]image.Rectangle) *SpriteSheet {
	return &SpriteSheet{
		Image: image,
		Slice: slice,
	}
}

// SliceImage creates a slice of image.Rectangle pointers
// that can be used to draw sprites from a sprite sheet.
func SliceImage(columns, rows, width, height int) [][]image.Rectangle {
	images := make([][]image.Rectangle, columns)

	for x := range images {
		images[x] = make([]image.Rectangle, rows)
		for y := range images[x] {
			images[x][y] = image.Rect(
				x*width,
				y*height,
				(x+1)*width,
				(y+1)*height,
			)
		}
	}

	return images
}

func (s *SpriteSheet) Sprite(x, y int) *ebiten.Image {
	return s.Image.SubImage(s.Slice[x][y]).(*ebiten.Image)
}
