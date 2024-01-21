package game

import (
	"fmt"
	"log/slog"
	"reflect"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/data"
	"github.com/MaikelVeen/go-game/ecs"
	"github.com/MaikelVeen/go-game/system"
	"github.com/MaikelVeen/go-game/timer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = &Game{}

type Game struct {
	config      *data.GameConfig
	coordinator *ecs.Coordinator
}

func New(config *data.GameConfig, coordinator *ecs.Coordinator) *Game {
	g := &Game{
		config:      config,
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
	defer timer.Timer("registerSystems")()

	// Register InputSystem.
	inputSystem := system.NewInputSystem(
		g.coordinator.ComponentManager,
	)
	g.coordinator.RegisterSystem(system.InputSystemType, inputSystem)
	g.coordinator.SetSystemSignature(system.InputSystemType, ecs.NewSignature(
		ecs.ComponentType(component.PlayerControllerType),
	))

	// Register PhysicsSystem.
	physicsSystem := system.NewPhysicsSystem(
		g.coordinator.ComponentManager,
	)
	g.coordinator.RegisterSystem(system.PhysicsSystemType, physicsSystem)
	g.coordinator.SetSystemSignature(system.PhysicsSystemType, ecs.NewSignature())

	// Register RenderSystem.
	renderSystem := system.NewRenderSystem(
		g.coordinator.ComponentManager,
		ebiten.NewImage(320, 240),
		4,
	)
	g.coordinator.RegisterSystem(system.RenderSystemType, renderSystem)
	g.coordinator.SetSystemSignature(system.RenderSystemType, ecs.NewSignature(
		ecs.ComponentType(component.TransformComponentType),
		ecs.ComponentType(component.SpriteRenderComponentType),
	))

	return nil
}

func (g *Game) createEntities() error {
	defer timer.Timer("createEntities")()

	// Create entities.
	for _, entityConfig := range g.config.Entities {
		entity, err := g.coordinator.CreateEntity()
		if err != nil {
			return err
		}

		// Add components to entity.
		for _, componentConfig := range entityConfig.Components {
			componentType, exists := component.ComponentMapping[componentConfig.Type]
			if !exists {
				return fmt.Errorf("unknown component type: %s", componentConfig.Type)
			}

			reflectType, exists := component.ComponentTypeMapping[componentType]
			if !exists {
				return fmt.Errorf("no reflect type found for component type: %d", componentType)
			}

			newComponentPtr := reflect.New(reflectType).Interface()
			newComponent, ok := newComponentPtr.(component.Component)
			if !ok {
				return fmt.Errorf("component does not implement Component interface: %s", componentConfig.Type)
			}

			if err := newComponent.SetData(componentConfig.Data); err != nil {
				return err
			}

			if err := g.coordinator.AddComponent(
				entity,
				ecs.ComponentType(componentType),
				newComponent,
			); err != nil {
				return err
			}
		}
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
