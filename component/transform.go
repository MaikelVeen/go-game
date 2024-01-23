package component

import (
	"fmt"

	"github.com/MaikelVeen/go-game/types"
)

const (
	TransformComponentName       = "transform"
	TransformComponentType uint8 = 0
)

var _ Component = (*Transform)(nil)

type Transform struct {
	Vector types.Vector2
}

// SetData implements Component.
func (t *Transform) SetData(data map[string]any) error {
	vec := types.Vector2{}

	x, ok := data["x"]
	if !ok {
		return fmt.Errorf("could not set transform data: missing x")
	}
	if x, ok := x.(float64); ok {
		vec.X = float64(x)
	}

	y, ok := data["y"]
	if !ok {
		return fmt.Errorf("could not set transform data: missing y")
	}
	if y, ok := y.(float64); ok {
		vec.Y = float64(y)
	}

	t.Vector = vec
	return nil
}
