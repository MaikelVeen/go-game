package system

import (
	"image/color"
	"strings"

	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const InputSystemType ecs.SystemType = 1

var _ ecs.System = &InputSystem{}

type InputSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}

	keys []ebiten.Key
}

func (s *InputSystem) AddEntity(entity ecs.Entity) {
	s.entities[entity] = struct{}{}
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

func (s *InputSystem) Update() error {
	s.keys = inpututil.AppendPressedKeys(s.keys[:0])
	return nil
}

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
