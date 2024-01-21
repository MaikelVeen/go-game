package system

import (
	"github.com/MaikelVeen/go-game/components"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const RenderSystemType ecs.SystemType = 1

var _ ecs.System = &RenderSystem{}

type RenderSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}
}

// There is a bug here if a system depends on multiple
func (s *RenderSystem) AddEntity(entity ecs.Entity) {
	s.entities[entity] = struct{}{}
}

func (s *RenderSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
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
	for entity := range s.entities {
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
