package game

import (
	"log/slog"

	"github.com/MaikelVeen/go-game/component"
	"github.com/MaikelVeen/go-game/entity"
	"github.com/MaikelVeen/go-game/system"
	"github.com/hajimehoshi/ebiten/v2"
)

// Coordinator is the main coordinator for the entity-component-system (ECS) architecture.
type Coordinator struct {
	EntityRegistry    *entity.Registry
	ComponentRegistry *component.Registry
	SystemRegistry    *system.Registry
}

func NewCoordinator(
	entityRegistry *entity.Registry,
	componentRegistry *component.Registry,
	systemRegistry *system.Registry,
) *Coordinator {
	return &Coordinator{
		EntityRegistry:    entityRegistry,
		ComponentRegistry: componentRegistry,
		SystemRegistry:    systemRegistry,
	}
}

func (c *Coordinator) Update() error {
	return c.SystemRegistry.ForEachSystem(func(sys system.System) error {
		return sys.Update()
	})
}

func (c *Coordinator) Draw(screen *ebiten.Image) {
	_ = c.SystemRegistry.ForEachSystem(func(sys system.System) error {
		sys.Draw(screen)
		return nil
	})
}

func (c *Coordinator) CreateEntity() (entity.Entity, error) {
	e, err := c.EntityRegistry.Create()
	slog.Debug("Creating entity", "entity", e)
	return e, err
}

func (c *Coordinator) DestroyEntity(entity entity.Entity) error {
	slog.Debug("Destroying entity", "entity", entity)
	if err := c.EntityRegistry.Destroy(entity); err != nil {
		return err
	}

	c.ComponentRegistry.EntityDestroyed(entity)
	c.SystemRegistry.EntityDestroyed(entity)
	return nil
}

func (c *Coordinator) AddComponent(entity entity.Entity, componentType component.Type, component any) error {
	slog.Debug("Adding component", "entity", entity, "componentType", componentType)

	if err := c.ComponentRegistry.AddComponent(entity, componentType, component); err != nil {
		return err
	}

	signature, err := c.EntityRegistry.Signature(entity)
	if err != nil {
		return err
	}
	signature.Set(uint(componentType))
	if err := c.EntityRegistry.SetSignature(entity, signature); err != nil {
		return err
	}

	c.SystemRegistry.EntitySignatureChanged(entity, signature)
	return nil
}

func (c *Coordinator) RemoveComponent(entity entity.Entity, componentType component.Type) error {
	slog.Debug("Removing component", "entity", entity, "componentType", componentType)

	if err := c.ComponentRegistry.RemoveComponent(entity, componentType); err != nil {
		return err
	}

	signature, err := c.EntityRegistry.Signature(entity)
	if err != nil {
		return err
	}
	signature.Set(uint(componentType))
	if err := c.EntityRegistry.SetSignature(entity, signature); err != nil {
		return err
	}

	c.SystemRegistry.EntitySignatureChanged(entity, signature)
	return nil
}

// RegisterSystem registers a system with the SystemManager.
// Returns an error if the system is already registered.
func (c *Coordinator) RegisterSystem(sysType system.Type, sys system.System) error {
	slog.Debug("Registering system", "systemType", sysType)
	return c.SystemRegistry.RegisterSystem(sysType, sys)
}

// SetSignature sets the signature for a system, this indicates which set of components
// the system is interested in.
func (c *Coordinator) SetSystemSignature(sysType system.Type, sig entity.Signature) error {
	slog.Debug("Setting system signature", "systemType", sysType, "signature", sig)
	return c.SystemRegistry.SetSignature(sysType, sig)
}
