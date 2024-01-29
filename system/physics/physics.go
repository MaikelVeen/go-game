package physics

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/component/rigidbody"
	"github.com/MaikelVeen/go-game/component/transform"
	"github.com/MaikelVeen/go-game/entity"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

const (
	Type uint8 = 1
	Slug       = "physics"
)

// PhysicsSystem is a system that handles physics.
type PhysicsSystem struct {
	componentRegistry *component.Registry
	entities          map[entity.Entity]*physicsEntity
	space             *cp.Space
}

type physicsEntity struct {
	rigidbody *rigidbody.Rigidbody
	transform *transform.Transform
}

// New returns a new PhysicsSystem.
func New(componentRegistry *component.Registry) *PhysicsSystem {
	return &PhysicsSystem{
		componentRegistry: componentRegistry,
		entities:          make(map[entity.Entity]*physicsEntity),
		space:             cp.NewSpace(),
	}
}

func (s *PhysicsSystem) AddEntity(entity entity.Entity) error {
	if _, exists := s.entities[entity]; exists {
		return nil
	}

	physicsEntity, err := s.initEntity(entity)
	if err != nil {
		return err
	}

	s.entities[entity] = physicsEntity
	s.space.AddBody(physicsEntity.rigidbody.Body)

	return nil
}

func (s *PhysicsSystem) initEntity(entity entity.Entity) (*physicsEntity, error) {
	t, err := s.componentRegistry.GetComponent(entity, component.TransformType)
	if err != nil {
		return nil, err
	}
	transform, ok := t.(*transform.Transform)
	if !ok {
		return nil, fmt.Errorf("could not typecast component to Transform")
	}

	rb, err := s.componentRegistry.GetComponent(entity, component.RigidbodyType)
	if err != nil {
		return nil, err
	}
	rigidbody, ok := rb.(*rigidbody.Rigidbody)
	if !ok {
		return nil, fmt.Errorf("could not typecast component to Rigidbody")
	}

	if err := rigidbody.Init(); err != nil {
		return nil, err
	}

	shape := cp.NewBox(rigidbody.Body, 16, 16, 0)
	rigidbody.Body.AddShape(shape)

	slog.Debug("Setting initial position of body", "entity", entity, "position", transform.Vector)
	rigidbody.Body.SetPosition(transform.Vector)

	s.space.AddShape(shape)
	s.space.AddBody(rigidbody.Body)

	return &physicsEntity{
		rigidbody: rigidbody,
		transform: transform,
	}, nil
}

func (s *PhysicsSystem) EntityDestroyed(entity entity.Entity) {
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
