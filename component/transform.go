package component

import "fmt"

const TransformComponentType uint8 = 0

var _ Component = (*Transform)(nil)

type Transform struct {
	X, Y int
}

func (t *Transform) SetData(data map[string]any) error {
	x, ok := data["x"]
	if !ok {
		return fmt.Errorf("could not set transform data: missing x")
	}
	if x, ok := x.(int); ok {
		t.X = x
	}

	y, ok := data["y"]
	if !ok {
		return fmt.Errorf("could not set transform data: missing y")
	}
	if y, ok := y.(int); ok {
		t.Y = y
	}

	return nil
}
