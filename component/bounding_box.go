package component

import (
	"github.com/MaikelVeen/go-game/types"
)

const (
	BoxColliderComponentName       = "boxCollider"
	BoxColliderComponentType uint8 = 3
)

var _ Component = (*BoundingBox)(nil)

type BoundingBox struct {
	Min *types.Vector2
	Max *types.Vector2
}

// SetData implements Component.
func (*BoundingBox) SetData(data map[string]any) error {
	// TODO: Implement.
	return nil
}
