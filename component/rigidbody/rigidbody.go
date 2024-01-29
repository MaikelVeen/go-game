package rigidbody

import (
	"fmt"

	"github.com/jakecoffman/cp/v2"
)

const (
	Type uint = 4
	Slug      = "rigidbody"
)

// RidigbodyPhysicsType is an enum for the type of rigidbody.
type RidigbodyPhysicsType uint8

const (
	// RigidbodyTypeDynamic is a rigidbody that moves and is affected by forces.
	// It collides with all other rigidbodies.
	RigidbodyTypeDynamic RidigbodyPhysicsType = iota
	// RigidbodyTypeKinematic is a rigidbody that moves and is affected by forces but
	// does not collide with other kinematic or static rigidbodies.
	RigidbodyTypeKinematic
	// RigidbodyTypeStatic is a rigidbody that does not move and is not affected by forces.
	// It collides with all other rigidbodies.
	RigidbodyTypeStatic
)

// RigidbodyTypeMapping maps strings to RidigbodyPhysicsType.
var RigidbodyTypeMapping = map[string]RidigbodyPhysicsType{
	"dynamic":   RigidbodyTypeDynamic,
	"kinematic": RigidbodyTypeKinematic,
	"static":    RigidbodyTypeStatic,
}

// Ridigbody wraps a Chipmunk body and implements PhysicsComponent.
// The body is created when Init is called.
type Rigidbody struct {
	Type RidigbodyPhysicsType
	*cp.Body
	mass *float64
}

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
