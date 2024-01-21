package input

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
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

// New returns a new InputSystem.
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

func (s *InputSystem) Init() error {
	return nil
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
//
// Each iteration of the game loop, relevant player controllers
// are updated with the current input direction and the rigidbody
// is updated with the new velocity.
func (s *InputSystem) Update() error {
	s.currentDirection = s.Direction()
	for entity := range s.entities {
		// Get the player controller component.
		pc, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.PlayerControllerType))
		if err != nil {
			return err
		}
		playerController := pc.(*component.PlayerController)

		_ = playerController.Speed

		// Get the rigidbody component.
		rb, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.RigidbodyComponentType))
		if err != nil {
			return err
		}
		_ = rb.(*component.Rigidbody)
	}

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
