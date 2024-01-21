package component

import "github.com/MaikelVeen/go-game/types"

const (
	RigidbodyComponentName       = "rigidbody"
	RigidbodyComponentType uint8 = 4
)

var _ Component = (*Rigidbody)(nil)

type Rigidbody struct {
	Velocity *types.Vector2
}

// SetData implements Component.
func (*Rigidbody) SetData(data map[string]any) error {
	return nil // Noop.
}

func (r *Rigidbody) AddForce(force *types.Vector2) {
	r.Velocity.Add(force)
}
