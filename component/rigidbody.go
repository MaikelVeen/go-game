package component

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

// TODO: Replace with iota in components.go.
const (
	RigidbodyComponentName       = "rigidbody"
	RigidbodyComponentType uint8 = 4
)

var _ PhysicsComponent = (*Rigidbody)(nil)

type RigidbodyType uint8

const (
	RigidbodyTypeDynamic RigidbodyType = iota
	RigidbodyTypeKinematic
	RigidbodyTypeStatic
)

var RigidbodyTypeMapping = map[string]RigidbodyType{
	"dynamic":   RigidbodyTypeDynamic,
	"kinematic": RigidbodyTypeKinematic,
	"static":    RigidbodyTypeStatic,
}

// Ridigbody wraps a Chipmunk body and implements PhysicsComponent.
// The body is created when Init is called.
type Rigidbody struct {
	*cp.Body
	Type RigidbodyType

	mass *float64
}

// Init implements PhysicsComponent.
func (r *Rigidbody) Init() error {
	switch r.Type {
	case RigidbodyTypeDynamic:
		if r.mass == nil {
			return fmt.Errorf("mass value not found")
		}

		r.Body = cp.NewBody(*r.mass, cp.INFINITY)
	case RigidbodyTypeKinematic:
		r.Body = cp.NewKinematicBody()
	case RigidbodyTypeStatic:
		r.Body = cp.NewStaticBody()
	}

	return nil
}

// SetData implements Component.
func (r *Rigidbody) SetData(data map[string]interface{}) error {
	massValue, massExists := data["mass"]
	if massExists {
		mass, ok := massValue.(float64)
		if !ok {
			return fmt.Errorf("mass value is not a float64")
		}
		r.mass = &mass
	}

	typeValue, typeExists := data["type"]
	if !typeExists {
		return fmt.Errorf("type value not found")
	}
	typ, ok := typeValue.(string)
	if !ok {
		return fmt.Errorf("type value is not a string")
	}
	t, ok := RigidbodyTypeMapping[typ]
	if !ok {
		return fmt.Errorf("invalid type value")
	}
	r.Type = t

	return nil
}
