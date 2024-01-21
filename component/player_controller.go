package component

const PlayerControllerComponentName = "playerController"
const PlayerControllerType uint8 = 2

var _ Component = (*PlayerController)(nil)

type PlayerController struct{}

// SetData implements Component.
func (*PlayerController) SetData(data map[string]any) error {
	return nil // Noop.
}
