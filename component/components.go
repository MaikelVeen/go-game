package component

import "reflect"

type Component interface {
	SetData(data map[string]any) error
}

var ComponentMapping = map[string]uint8{
	"transform":        TransformComponentType,
	"spriteRender":     SpriteRenderComponentType,
	"playerController": PlayerControllerType,
}

var ComponentTypeMapping = map[uint8]reflect.Type{
	TransformComponentType:    reflect.TypeOf(Transform{}),
	PlayerControllerType:      reflect.TypeOf(PlayerController{}),
	SpriteRenderComponentType: reflect.TypeOf(SpriteRender{}),
}
