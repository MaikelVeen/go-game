package system

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

const RenderSystemType ecs.SystemType = 2

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

func (s *RenderSystem) AddEntity(entity ecs.Entity) {
	if _, exists := s.entities[entity]; exists {
		return
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to RenderSystem", "entity", entity)
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
		drawEntity(s, entity)
	}

	// Scale the offscreen image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(s.scaleFactor, s.scaleFactor)

	screen.DrawImage(s.offScreenImage, op)
}

func drawEntity(s *RenderSystem, entity ecs.Entity) {
	t, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.TransformComponentType))
	if err != nil {
		slog.Error("Failed to get Transform component", "entity", entity)
		return
	}
	tranform := t.(*component.Transform)

	spriteRender, err := s.componentManager.GetComponent(entity, ecs.ComponentType(component.SpriteRenderComponentType))
	if err != nil {
		slog.Error(err.Error())
		return
	}
	sr := spriteRender.(*component.SpriteRender)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		tranform.Vector.X,
		tranform.Vector.Y,
	)

	s.offScreenImage.DrawImage(sr.GetSprite(), op)
}
