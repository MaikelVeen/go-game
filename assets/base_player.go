package assets

import "github.com/hajimehoshi/ebiten/v2"

var (
	PlayerSpriteSheet = SliceImage(3, 4, 16, 16)
	PlayerIdle        *ebiten.Image
)

func init() {
	img, err := loadFile(staticSpritesFS, "sprites/base_player.png")
	if err != nil {
		panic(err)
	}

	ebitenImage := ebiten.NewImageFromImage(img)
	PlayerIdle = ebitenImage.SubImage(PlayerSpriteSheet[0][1]).(*ebiten.Image)
}