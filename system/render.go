package system

import (
	"github.com/MaikelVeen/go-game/components"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const RenderSystemType ecs.SystemType = 0

var _ ecs.System = &RenderSystem{}

type RenderSystem struct {
	componentManager *ecs.ComponentManager
	entities         map[ecs.Entity]struct{}

	// offScreenImage is the image that is rendered to by the RenderSystem.
	// This image is then scaled and drawn to the screen.
	offScreenImage *ebiten.Image

	// scaleFactor is the factor by which the offScreenImage is scaled.
	scaleFactor float64
}

// There is a bug here if a system depends on multiple
func (s *RenderSystem) AddEntity(entity ecs.Entity) {
	s.entities[entity] = struct{}{}
}

func (s *RenderSystem) EntityDestroyed(entity ecs.Entity) {
	delete(s.entities, entity)
}

func NewRenderSystem(
	cm *ecs.ComponentManager,
	offScreenImage *ebiten.Image,
	scale float64,
) *RenderSystem {
	return &RenderSystem{
		componentManager: cm,
		entities:         make(map[ecs.Entity]struct{}),
		offScreenImage:   offScreenImage,
		scaleFactor:      scale,
	}
}

func (s *RenderSystem) Update() error {
	return nil // Noop.
}

func (s *RenderSystem) Draw(screen *ebiten.Image) {
	s.offScreenImage.Clear()

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

		s.offScreenImage.DrawImage(sr.Image, op)
	}

	// Scale the offscreen image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(s.scaleFactor), float64(s.scaleFactor))

	screen.DrawImage(s.offScreenImage, op)
}
