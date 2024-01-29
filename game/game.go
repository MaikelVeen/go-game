package game

import (
	"fmt"
	"reflect"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/component/boxcollider"
	"github.com/MaikelVeen/go-game/component/playercontroller"
	"github.com/MaikelVeen/go-game/component/rigidbody"
	"github.com/MaikelVeen/go-game/component/spriterenderer"
	"github.com/MaikelVeen/go-game/component/transform"
	"github.com/MaikelVeen/go-game/data"
	"github.com/MaikelVeen/go-game/entity"
	"github.com/MaikelVeen/go-game/system"
	"github.com/MaikelVeen/go-game/system/input"
	"github.com/MaikelVeen/go-game/system/physics"
	"github.com/MaikelVeen/go-game/system/render"
	"github.com/MaikelVeen/go-game/timer"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = &Game{}

type Game struct {
	config      *data.GameConfig
	coordinator *Coordinator
}

// New returns a new Game.
func New(config *data.GameConfig, coordinator *Coordinator) (*Game, error) {
	g := &Game{
		config:      config,
		coordinator: coordinator,
	}

	if err := g.registerSystems(); err != nil {
		return nil, err
	}

	if err := g.createEntities(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) registerSystems() error {
	defer timer.Timer("registerSystems")()

	// Register InputSystem.
	if err := g.registerSystem(
		system.InputType,
		input.New(g.coordinator.ComponentRegistry),
		entity.NewSignature(
			playercontroller.Type,
		),
	); err != nil {
		return err
	}

	// Register PhysicsSystem.
	if err := g.registerSystem(
		system.PhysicsType,
		physics.New(g.coordinator.ComponentRegistry),
		entity.NewSignature(
			transform.Type,
			rigidbody.Type,
			boxcollider.Type, // When there are multiple colliders, how should this be handled?
		),
	); err != nil {
		return err
	}

	// Register RenderSystem.
	if err := g.registerSystem(
		system.RenderType,
		render.New(g.coordinator.ComponentRegistry, ebiten.NewImage(640, 330)),
		entity.NewSignature(
			transform.Type,
			spriterenderer.Type),
	); err != nil {
		return err
	}

	return nil
}

// registerSystem registers a system and sets its signature.
func (g *Game) registerSystem(sysType system.Type, sys system.System, sig entity.Signature) error {
	if err := g.coordinator.RegisterSystem(sysType, sys); err != nil {
		return fmt.Errorf("failed to register system: %w", err)
	}
	if err := g.coordinator.SetSystemSignature(sysType, sig); err != nil {
		return fmt.Errorf("failed to set system signature: %w", err)
	}
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
				componentType,
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
