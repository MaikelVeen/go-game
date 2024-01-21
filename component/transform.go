package component

const TransformComponentType uint8 = 0

var _ Component = (*Transform)(nil)

type Transform struct {
	X, Y int
}

func (*Transform) SetData(data map[string]any) error {
	return nil // TODO: Implement.
}
