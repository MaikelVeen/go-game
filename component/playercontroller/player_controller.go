package playercontroller

import (
	"fmt"
)

const (
	Type uint = 2
	Slug      = "playerController"
)

type PlayerController struct {
	Speed float64
}

func (pc *PlayerController) SetData(data map[string]any) error {
	if speed, ok := data["speed"].(float64); ok {
		pc.Speed = speed
	} else {
		return fmt.Errorf("invalid speed: %v", data["speed"])
	}

	return nil
}
