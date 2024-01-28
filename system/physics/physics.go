package physics

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/ecs"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

const SystemType ecs.SystemType = 1

var _ ecs.System = &PhysicsSystem{}

// PhysicsSystem is a system that handles physics.
type PhysicsSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct {
		rigidbody *component.Rigidbody
		transform *component.Transform
	}
	space *cp.Space
}

// New returns a new PhysicsSystem.
func New(
	cm *ecs.ComponentManager,
	space *cp.Space,
) *PhysicsSystem {
	return &PhysicsSystem{
		componentManager: cm,
		entities: make(map[ecs.Entity]struct {
			rigidbody *component.Rigidbody
			transform *component.Transform
		}),
		space: space,
	}
}

func (s *PhysicsSystem) AddEntity(entity ecs.Entity) error {
	if _, exists := s.entities[entity]; exists {
		return nil
	}

	components, err := s.initEntity(entity)
	if err != nil {
		return err
	}

	s.entities[entity] = *components
	s.space.AddBody(components.rigidbody.Body)
	slog.Debug("Added entity to InputSystem", "entity", entity)
	return nil
}

func (s *PhysicsSystem) initEntity(entity ecs.Entity) (*struct {
	rigidbody *component.Rigidbody
	transform *component.Transform
}, error) {
	t, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.TransformComponentType))
	if err != nil {
		return nil, err
	}
	transform, ok := t.(*component.Transform)
	if !ok {
		return nil, fmt.Errorf("could not typecast component to *component.Transform")
	}

	rb, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.RigidbodyComponentType))
	if err != nil {
		return nil, err
	}
	rigidbody, ok := rb.(*component.Rigidbody)
	if !ok {
		return nil, fmt.Errorf("could not typecast component to *component.Rigidbody")
	}
	//TODO: Add a test to assert that this function returns an error when the mass is not set and the type is dynamic.
	if err := rigidbody.Init(); err != nil {
		return nil, err
	}

	// TODO: Collision callbacks to ensure that the player does not glitch into walls.
	shape := cp.NewBox(rigidbody.Body, 16, 16, 0)
	rigidbody.Body.AddShape(shape)

	slog.Debug("Setting initial position of body", "entity", entity, "position", transform.Vector)
	rigidbody.Body.SetPosition(transform.Vector)

	s.space.AddShape(shape)
	return &struct {
		rigidbody *component.Rigidbody
		transform *component.Transform
	}{
		rigidbody: rigidbody,
		transform: transform,
	}, nil
}

func (s *PhysicsSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
}

// Draw implements ecs.System.
func (*PhysicsSystem) Draw(screen *ebiten.Image) {} // Noop.

// Update implements ecs.System.
func (s *PhysicsSystem) Update() error {
	s.space.Step(1.0 / float64(ebiten.TPS()))

	for _, components := range s.entities {
		components.transform.Vector = components.rigidbody.Body.Position()
	}
	return nil
}
