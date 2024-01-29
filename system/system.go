package system

import (
	"github.com/MaikelVeen/go-game/entity"
	"github.com/MaikelVeen/go-game/system/input"
	"github.com/MaikelVeen/go-game/system/physics"
	"github.com/MaikelVeen/go-game/system/render"
	"github.com/hajimehoshi/ebiten/v2"
)

// Type is an unique identifier for a system.
type Type uint8

const (
	InputType   Type = Type(input.Type)
	PhysicsType Type = Type(physics.Type)
	RenderType  Type = Type(render.Type)
)

var (
	_ System = (*input.InputSystem)(nil)
	_ System = (*physics.PhysicsSystem)(nil)
	_ System = (*render.RenderSystem)(nil)
)

type System interface {
	// TODO: Extract this to a seperate interface?.
	AddEntity(entity.Entity) error
	EntityDestroyed(entity.Entity)

	Update() error
	Draw(*ebiten.Image)
}
