package system

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/ecs"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

const PhysicsSystemType ecs.SystemType = 1

var _ ecs.System = &PhysicsSystem{}

// PhysicsSystem is a system that handles physics.
type PhysicsSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}
}

// NewPhysicsSystem returns a new PhysicsSystem.
func NewPhysicsSystem(cm *ecs.ComponentManager) *PhysicsSystem {
	return &PhysicsSystem{
		componentManager: cm,
		entities:         make(map[ecs.Entity]struct{}),
	}
}

func (s *PhysicsSystem) AddEntity(entity ecs.Entity) {
	// Check if exists.
	if _, exists := s.entities[entity]; exists {
		return
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to InputSystem", "entity", entity)
}

func (s *PhysicsSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
}

// Draw implements ecs.System.
func (*PhysicsSystem) Draw(screen *ebiten.Image) {} // Noop.

// Update implements ecs.System.
func (*PhysicsSystem) Update() error {
	// TODO: Implement.
	return nil
}
