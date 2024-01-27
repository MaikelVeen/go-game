package component

import (
	"github.com/MaikelVeen/go-game/types"
)

const (
	BoxColliderComponentName       = "boxCollider"
	BoxColliderComponentType uint8 = 3
)

var _ Component = (*BoxCollider)(nil)

type BoxCollider struct {
	Min *types.Vector2
	Max *types.Vector2
}

// SetData implements Component.
func (*BoxCollider) SetData(data map[string]any) error {
	// TODO: Implement.
	return nil
}
