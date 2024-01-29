package render

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/component/spriterenderer"
	"github.com/MaikelVeen/go-game/component/transform"
	"github.com/MaikelVeen/go-game/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Type uint8 = 2
	Slug       = "render"
)

type RenderSystem struct {
	componentRegistry *component.Registry

	entities map[entity.Entity]struct{}

	// offScreenImage is the image that is rendered to by the RenderSystem.
	// This image is then scaled and drawn to the screen.
	offScreenImage *ebiten.Image
}

// New returns a new RenderSystem.
func New(componentRegistry *component.Registry, offscreenImage *ebiten.Image) *RenderSystem {
	return &RenderSystem{
		componentRegistry: componentRegistry,
		offScreenImage:    offscreenImage,
		entities:          make(map[entity.Entity]struct{}),
	}
}

func (s *RenderSystem) Init() error {
	return nil
}

func (s *RenderSystem) AddEntity(entity entity.Entity) error {
	if _, exists := s.entities[entity]; exists {
		return nil
	}

	s.entities[entity] = struct{}{}
	slog.Debug("Added entity to RenderSystem", "entity", entity)
	return nil
}

func (s *RenderSystem) EntityDestroyed(entity entity.Entity) {
	delete(s.entities, entity)
}

func (s *RenderSystem) Update() error {
	return nil // Noop.
}

func (s *RenderSystem) Draw(screen *ebiten.Image) {
	s.offScreenImage.Clear()

	for entity := range s.entities {
		drawEntity(s, entity)
	}

	screenWidth, screenHeight := getScreenSize()
	scaleFactorX := float64(screenWidth) / float64(s.offScreenImage.Bounds().Dx())
	scaleFactorY := float64(screenHeight) / float64(s.offScreenImage.Bounds().Dy())

	// Scale the offscreen image.
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleFactorX, scaleFactorY)
	screen.DrawImage(s.offScreenImage, op)
}

func getScreenSize() (int, int) {
	if ebiten.IsFullscreen() {
		return ebiten.ScreenSizeInFullscreen()
	}
	return ebiten.WindowSize()
}

// TODO: Make this part of the RenderSystem.
func drawEntity(s *RenderSystem, entity entity.Entity) {
	t, err := s.componentRegistry.GetComponent(entity, component.TransformType)
	if err != nil {
		slog.Error("Failed to get Transform component", "entity", entity)
		return
	}
	tranform := t.(*transform.Transform)

	sr, err := s.componentRegistry.GetComponent(entity, component.SpriteRenderType)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	spriteRenderer, ok := sr.(*spriterenderer.SpriteRenderer)
	if !ok {
		slog.Error("Failed to typecast SpriteRenderer component")
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		tranform.Vector.X,
		tranform.Vector.Y,
	)

	s.offScreenImage.DrawImage(spriteRenderer.GetSprite(), op)
}
