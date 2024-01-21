package system

import (
	"github.com/MaikelVeen/go-game/components"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

var _ ecs.System = &RenderSystem{}

type RenderSystem struct {
	componentManager *ecs.ComponentManager
	entities         []ecs.Entity
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

func NewRenderSystem(cm *ecs.ComponentManager) *RenderSystem {
	return &RenderSystem{
		componentManager: cm,
	}
}

func (s *RenderSystem) Update() error {
	return nil // Noop.
}

func (s *RenderSystem) Draw(screen *ebiten.Image) {
	println("RenderSystem.Draw")
	for _, entity := range s.entities {
		t, err := s.componentManager.GetComponent(entity, ecs.ComponentType(components.TransformComponentType))
		if err != nil {
			// TODO : Log an error here.
			panic(err)
		}
		tranform := t.(*components.Transform)

		spriteRender, err := s.componentManager.GetComponent(entity, ecs.ComponentType(components.SpriteRenderComponentType))
		if err != nil {
			// TODO : Log an error here.
			panic(err)
		}
		sr := spriteRender.(*components.SpriteRender)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(
			float64(tranform.X),
			float64(tranform.Y),
		)

		screen.DrawImage(sr.Image, op)
	}
}
