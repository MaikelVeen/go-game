package system

import (
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ ecs.System = &RenderSystem{}

type RenderSystem struct {
	screen   *ebiten.Image
	entities []ecs.Entity
}

func (s *RenderSystem) AddEntity(entity ecs.Entity) {
	s.entities = append(s.entities, entity)
}

func (s *RenderSystem) EntityDestroyed(entity ecs.Entity) {
	for i, e := range s.entities {
		if e == entity {
			s.entities = append(s.entities[:i], s.entities[i+1:]...)
			break
		}
	}
}

func NewRenderSystem(screen *ebiten.Image) *RenderSystem {
	return &RenderSystem{
		screen: screen,
	}
}

func (s *RenderSystem) Update() {

}
