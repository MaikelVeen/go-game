package component

import "reflect"

type Component interface {
	SetData(data map[string]any) error
}

// PhysicsComponent is a component that is used by the physics system.
type PhysicsComponent interface {
	Component
	Init() error
}

// ComponentMapping maps a component name to a component type.
// In the data.GameConfig struct, the component names are used
// for the components of entities.
var ComponentMapping = map[string]uint8{
	TransformComponentName:        TransformComponentType,
	SpriteRenderComponentName:     SpriteRenderComponentType,
	PlayerControllerComponentName: PlayerControllerType,
	BoxColliderComponentName:      BoxColliderComponentType,
	RigidbodyComponentName:        RigidbodyComponentType,
}

// ComponentTypeMapping maps a component type to a component reflect.Type.
// This is used to create new components of a certain type during startup.
var ComponentTypeMapping = map[uint8]reflect.Type{
	TransformComponentType:    reflect.TypeOf(Transform{}),
	PlayerControllerType:      reflect.TypeOf(PlayerController{}),
	SpriteRenderComponentType: reflect.TypeOf(SpriteRender{}),
	BoxColliderComponentType:  reflect.TypeOf(BoundingBox{}),
	RigidbodyComponentType:    reflect.TypeOf(Rigidbody{}),
}
