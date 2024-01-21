package physics

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/ecs"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"
)

const SystemType ecs.SystemType = 1

var _ ecs.System = &PhysicsSystem{}

// PhysicsSystem is a system that handles physics.
type PhysicsSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}
	space            *cp.Space
}

// New returns a new PhysicsSystem.
func New(
	cm *ecs.ComponentManager,
	space *cp.Space,
) *PhysicsSystem {
	return &PhysicsSystem{
		componentManager: cm,
		entities:         make(map[ecs.Entity]struct{}),
		space:            space,
	}
}

// Init implements ecs.System.
func (s *PhysicsSystem) Init() error {
	return nil
}

func (s *PhysicsSystem) AddEntity(entity ecs.Entity) {
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