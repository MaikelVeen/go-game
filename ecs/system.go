package ecs

import "github.com/hajimehoshi/ebiten/v2"

// ComponentType is a unique identifier for a type of component.
type SystemType uint8

type System interface {
	// TODO: Extract this to a seperate interface?.
	AddEntity(entity Entity)
	EntityDestroyed(entity Entity)

	Update() error
	Draw(screen *ebiten.Image)
}
