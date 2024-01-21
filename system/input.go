package system

import (
	"image/color"
	"log/slog"
	"strings"

	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const InputSystemType ecs.SystemType = 0

var _ ecs.System = &InputSystem{}

type InputSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}

	keys []ebiten.Key
}

func (s *InputSystem) AddEntity(entity ecs.Entity) {
	// Check if exists.
	if _, exists := s.entities[entity]; exists {
		return
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to InputSystem", "entity", entity)
}

func (s *InputSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
}

func NewInputSystem(cm *ecs.ComponentManager) *InputSystem {
	return &InputSystem{
		componentManager: cm,
		entities:         make(map[ecs.Entity]struct{}),
	}
}

// Draw implements ecs.System.
func (s *InputSystem) Draw(screen *ebiten.Image) {
	var keyStrs []string
	var keyNames []string
	for _, k := range s.keys {
		keyStrs = append(keyStrs, k.String())
		if name := ebiten.KeyName(k); name != "" {
			keyNames = append(keyNames, name)
		}
	}

	text.Draw(
		screen,
		strings.Join(keyStrs, ", ")+"\n"+strings.Join(keyNames, ", "),
		bitmapfont.Face,
		0,
		40,
		color.White,
	)
}

// Update implements ecs.System.
func (s *InputSystem) Update() error {
	s.keys = inpututil.AppendPressedKeys(s.keys[:0])
	return nil
}
