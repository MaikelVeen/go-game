package component

import (
	"github.com/jakecoffman/cp/v2"
)

const (
	RigidbodyComponentName       = "rigidbody"
	RigidbodyComponentType uint8 = 4
)

var _ PhysicsComponent = (*Rigidbody)(nil)

type Rigidbody struct {
	*cp.Body

	Mass   float64
	Static bool
}

// Init implements PhysicsComponent.
func (r *Rigidbody) Init() error {
	if r.Static {
		r.Body = cp.NewStaticBody()
		return nil
	}

	r.Body = cp.NewBody(r.Mass, cp.INFINITY)
	return nil
}

// SetData implements Component.
func (r *Rigidbody) SetData(data map[string]any) error {
	if mass, ok := data["mass"].(float64); ok {
		r.Mass = mass
	}

	if static, ok := data["static"].(bool); ok {
		r.Static = static
	}

	return nil
}
