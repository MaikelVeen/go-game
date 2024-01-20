package game

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/components"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = &Game{}

// Game is the main game struct.
// It encapsulates the entity-component-system (ECS) architecture.
type Game struct {
	entityManager    *ecs.EntityManager
	componentManager *ecs.ComponentManager
	systemManager    *ecs.SystemManager

	systems []ecs.System
}

// TODO: New Should take a list of dependencies and logger.
func New() *Game {
	g := &Game{
		entityManager:    ecs.NewEntityManager(),
		componentManager: ecs.NewComponentManager(),
		systemManager:    ecs.NewSystemManager(),
	}

	if err := g.registerSystems(); err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if err := g.createEntities(); err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	return g
}

func (g *Game) Init() *Game {
	// TODO: Add init logic here for systems.
	return g
}

func (g *Game) registerSystems() error {
	// Register RenderSystem.
	renderSystem := system.NewRenderSystem(g.componentManager)
	g.systemManager.RegisterSystem(renderSystem)

	// Set signature for RenderSystem.
	signature := ecs.NewSignature(
		ecs.ComponentType(components.TransformComponentType),
		ecs.ComponentType(components.SpriteRenderComponentType),
	)
	g.systemManager.SetSignature(renderSystem, signature)

	return nil
}

func (g *Game) createEntities() error {
	player, err := g.entityManager.Create()
	if err != nil {
		return err
	}

	if err := g.componentManager.AddComponent(
		player,
		ecs.ComponentType(components.TransformComponentType),
		&components.Transform{
			X: 10, Y: 10,
		},
	); err != nil {
		return err
	}

	if err := g.componentManager.AddComponent(
		player,
		ecs.ComponentType(components.SpriteRenderComponentType),
		&components.SpriteRender{},
	); err != nil {
		return err
	}

	return nil
}

func (g *Game) Update() error {
	for _, system := range g.systems {
		if err := system.Update(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))

	for _, system := range g.systems {
		system.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
