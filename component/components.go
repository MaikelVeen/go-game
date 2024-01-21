package component

import "reflect"

type Component interface {
	SetData(data map[string]any) error
}

// ComponentMapping maps a component name to a component type.
// In the data.GameConfig struct, the component names are used
// for the components of entities.
var ComponentMapping = map[string]uint8{
	TransformComponentName:        TransformComponentType,
	SpriteRenderComponentName:     SpriteRenderComponentType,
	PlayerControllerComponentName: PlayerControllerType,
	BoundingBoxComponentName:      BoundingBoxComponentType,
	RigidbodyComponentName:        RigidbodyComponentType,
}

// ComponentTypeMapping maps a component type to a component reflect.Type.
// This is used to create new components of a certain type during startup.
var ComponentTypeMapping = map[uint8]reflect.Type{
	TransformComponentType:    reflect.TypeOf(Transform{}),
	PlayerControllerType:      reflect.TypeOf(PlayerController{}),
	SpriteRenderComponentType: reflect.TypeOf(SpriteRender{}),
	BoundingBoxComponentType:  reflect.TypeOf(BoundingBox{}),
	RigidbodyComponentType:    reflect.TypeOf(Rigidbody{}),
}
