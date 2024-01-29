package component

import (
	"reflect"

	"github.com/MaikelVeen/go-game/component/boxcollider"
	"github.com/MaikelVeen/go-game/component/playercontroller"
	"github.com/MaikelVeen/go-game/component/rigidbody"
	"github.com/MaikelVeen/go-game/component/spriterenderer"
	"github.com/MaikelVeen/go-game/component/transform"
)

// Type is an unique identifier for a type of component.
type Type uint8

const (
	TransformType        Type = Type(transform.Type)
	SpriteRenderType     Type = Type(spriterenderer.Type)
	PlayerControllerType Type = Type(playercontroller.Type)
	BoxColliderType      Type = Type(boxcollider.Type)
	RigidbodyType        Type = Type(rigidbody.Type)
)

var (
	_ Component = (*transform.Transform)(nil)
	_ Component = (*spriterenderer.SpriteRenderer)(nil)
	_ Component = (*playercontroller.PlayerController)(nil)
	_ Component = (*boxcollider.BoxCollider)(nil)
	_ Component = (*rigidbody.Rigidbody)(nil)
)

// Component is a component that can be added to an entity.
type Component interface {
	SetData(data map[string]any) error
}

// TODO: Think about if this is needed or Init can be merged with SetData.
// PhysicsComponent is a component that is used by the physics system.
type PhysicsComponent interface {
	Component
	Init() error
}

// ComponentMapping maps a component name to a component type.
//
// In the data.GameConfig struct, the component names are used
// for the components of entities.
var ComponentMapping = map[string]Type{
	transform.Slug:        TransformType,
	spriterenderer.Slug:   SpriteRenderType,
	playercontroller.Slug: PlayerControllerType,
	boxcollider.Slug:      BoxColliderType,
	rigidbody.Slug:        RigidbodyType,
}

// ComponentTypeMapping maps a component type to a component reflect.Type.
//
// This is used to create new components of a certain type during startup.
var ComponentTypeMapping = map[Type]reflect.Type{
	TransformType:        reflect.TypeOf(transform.Transform{}),
	SpriteRenderType:     reflect.TypeOf(spriterenderer.SpriteRenderer{}),
	PlayerControllerType: reflect.TypeOf(playercontroller.PlayerController{}),
	BoxColliderType:      reflect.TypeOf(boxcollider.BoxCollider{}),
	RigidbodyType:        reflect.TypeOf(rigidbody.Rigidbody{}),
}
