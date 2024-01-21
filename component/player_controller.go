package component

import (
	"fmt"

	"github.com/MaikelVeen/go-game/types"
)

const PlayerControllerComponentName = "playerController"
const PlayerControllerType uint8 = 2

var _ Component = (*PlayerController)(nil)

type PlayerController struct {
	Speed     float64
	Direction *types.Vector2
}

func (pc *PlayerController) Update(Direction *types.Vector2) error {
	pc.Direction = Direction
	return nil
}

// SetData implements Component.
func (pc *PlayerController) SetData(data map[string]any) error {
	// Set the speed.
	if speed, ok := data["speed"].(float64); ok {
		pc.Speed = speed
	} else {
		return fmt.Errorf("invalid speed: %v", data["speed"])
	}

	return nil
}
