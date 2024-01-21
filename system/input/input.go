package input

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const SystemType ecs.SystemType = 0

var _ ecs.System = &InputSystem{}

type InputSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}

	currentDirection *types.Vector2
}

func New(cm *ecs.ComponentManager) *InputSystem {
	return &InputSystem{
		componentManager: cm,
		entities:         make(map[ecs.Entity]struct{}),
		currentDirection: &types.Vector2{},
	}
}

func (s *InputSystem) AddEntity(entity ecs.Entity) {
	if _, exists := s.entities[entity]; exists {
		return
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to InputSystem", "entity", entity)
}

func (s *InputSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
}

// Draw implements ecs.System.
func (s *InputSystem) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"Input direction X: %f, Y: %f",
			s.currentDirection.X,
			s.currentDirection.Y),
		0,
		15,
	)
}

// Update implements ecs.System.
func (s *InputSystem) Update() error {
	s.currentDirection = s.Direction()
	return nil
}

// Direction returns the direction of the input.
func (s *InputSystem) Direction() *types.Vector2 {
	var direction types.Vector2

	if UpKeys.PressedAny() {
		direction.Y = -1
	}

	if DownKeys.PressedAny() {
		direction.Y = 1
	}

	if LeftKeys.PressedAny() {
		direction.X = -1
	}

	if RightKeys.PressedAny() {
		direction.X = 1
	}

	return &direction
}
