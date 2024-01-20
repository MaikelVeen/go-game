package ecs

import "github.com/hajimehoshi/ebiten/v2"

type System interface {
	AddEntity(entity Entity)
	EntityDestroyed(entity Entity)

	Update() error
	Draw(screen *ebiten.Image)
}
