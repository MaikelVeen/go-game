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

func (s *InputSystem) AddEntity(entity ecs.Entity) error {
	if _, exists := s.entities[entity]; exists {
		return nil
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to InputSystem", "entity", entity)
	return nil
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
		playerController, ok := pc.(*component.PlayerController)
		if !ok {
			return fmt.Errorf("could not typecast component to PlayerController")
		}
		speed := playerController.Speed

		// TOOD: Should the input system be responsible for updatring the velocity of the rigidbody?
		// Get the rigidbody component.
		rb, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.RigidbodyComponentType))
		if err != nil {
			return err
		}
		rigidbody, ok := rb.(*component.Rigidbody)
		if !ok {
			return fmt.Errorf("could not typecast component to Rigidbody")
		}

		if s.currentDirection.Equal(types.Vector2{}) {
			rigidbody.Body.SetVelocity(0, 0)
			continue
		}

		// Normalize the direction and multiply it by the speed.
		velocity := s.currentDirection.Normalize().Mult(speed)
		rigidbody.Body.SetForce(velocity)

		// Apply drag to the current velocity.
		dragFactor := 0.9
		currentVelocity := rigidbody.Body.Velocity()
		draggedVelocity := currentVelocity.Mult(dragFactor)

		// Set the new velocity of the rigidbody.
		rigidbody.Body.SetVelocityVector(draggedVelocity)
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
