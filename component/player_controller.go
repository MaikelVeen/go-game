package component

const PlayerControllerType uint8 = 2

var _ Component = (*PlayerController)(nil)

type PlayerController struct{}

func (*PlayerController) SetData(data map[string]any) error {
	return nil // Noop.
}
