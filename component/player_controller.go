package component

import (
	"fmt"
)

const PlayerControllerComponentName = "playerController"
const PlayerControllerType uint8 = 2

var _ Component = (*PlayerController)(nil)

type PlayerController struct {
	Speed float64
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
