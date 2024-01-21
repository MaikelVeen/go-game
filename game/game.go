package game

import (
	"fmt"
	"log/slog"

	"github.com/MaikelVeen/go-game/assets"
	"github.com/MaikelVeen/go-game/components"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/system"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = &Game{}

type Game struct {
	coordinator *ecs.Coordinator
}

// TODO: Inject Logger.
func New(coordinator *ecs.Coordinator) *Game {
	g := &Game{
		coordinator: coordinator,
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

func (g *Game) registerSystems() error {
	// Register RenderSystem.
	renderSystem := system.NewRenderSystem(
		g.coordinator.ComponentManager,
		ebiten.NewImage(320, 240),
		4,
	)
	g.coordinator.RegisterSystem(system.RenderSystemType, renderSystem)

	// Set signature for RenderSystem.
	signature := ecs.NewSignature(
		ecs.ComponentType(components.TransformComponentType),
		ecs.ComponentType(components.SpriteRenderComponentType),
	)
	g.coordinator.SetSystemSignature(system.RenderSystemType, signature)

	// Register InputSystem.
	inputSystem := system.NewInputSystem(
		g.coordinator.ComponentManager,
	)
	g.coordinator.RegisterSystem(system.InputSystemType, inputSystem)

	// Set signature for InputSystem.
	signature = ecs.NewSignature(
		ecs.ComponentType(components.PlayerControllerType),
	)
	g.coordinator.SetSystemSignature(system.InputSystemType, signature)

	return nil
}

func (g *Game) createEntities() error {
	player, err := g.coordinator.CreateEntity()
	if err != nil {
		return err
	}

	if err := g.coordinator.AddComponent(
		player,
		ecs.ComponentType(components.TransformComponentType),
		&components.Transform{
			X: 100, Y: 100,
		},
	); err != nil {
		return err
	}

	if err := g.coordinator.AddComponent(
		player,
		ecs.ComponentType(components.SpriteRenderComponentType),
		&components.SpriteRender{
			Image: assets.PlayerIdle,
		},
	); err != nil {
		return err
	}

	return nil
}

func (g *Game) Update() error {
	return g.coordinator.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	g.coordinator.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
