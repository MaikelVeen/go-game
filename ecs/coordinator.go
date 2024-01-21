package ecs

import "github.com/hajimehoshi/ebiten/v2"

// Coordinator is the main coordinator for the entity-component-system (ECS) architecture.
type Coordinator struct {
	EntityManager    *EntityManager
	ComponentManager *ComponentManager
	SystemManager    *SystemManager
}

func NewCoordinator(
	entityManager *EntityManager,
	componentManager *ComponentManager,
	systemManager *SystemManager,
) *Coordinator {
	return &Coordinator{
		EntityManager:    entityManager,
		ComponentManager: componentManager,
		SystemManager:    systemManager,
	}
}

func (c *Coordinator) Update() error {
	return c.SystemManager.ForEachSystem(func(sys System) error {
		return sys.Update()
	})
}

func (c *Coordinator) Draw(screen *ebiten.Image) {
	_ = c.SystemManager.ForEachSystem(func(sys System) error {
		sys.Draw(screen)
		return nil
	})
}

func (c *Coordinator) CreateEntity() (Entity, error) {
	return c.EntityManager.Create()
}

func (c *Coordinator) DestroyEntity(entity Entity) error {
	return c.EntityManager.Destroy(entity)
}

func (c *Coordinator) AddComponent(entity Entity, componentType ComponentType, component any) error {
	if err := c.ComponentManager.AddComponent(entity, componentType, component); err != nil {
		return err
	}

	signature, err := c.EntityManager.Signature(entity)
	if err != nil {
		return err
	}
	signature.Set(uint(componentType))
	if err := c.EntityManager.SetSignature(entity, signature); err != nil {
		return err
	}

	c.SystemManager.EntitySignatureChanged(entity, signature)
	return nil
}

func (c *Coordinator) RemoveComponent(entity Entity, componentType ComponentType) error {
	if err := c.ComponentManager.RemoveComponent(entity, componentType); err != nil {
		return err
	}

	signature, err := c.EntityManager.Signature(entity)
	if err != nil {
		return err
	}
	signature.Set(uint(componentType))
	if err := c.EntityManager.SetSignature(entity, signature); err != nil {
		return err
	}

	c.SystemManager.EntitySignatureChanged(entity, signature)
	return nil
}

// RegisterSystem registers a system with the SystemManager.
// Returns an error if the system is already registered.
func (c *Coordinator) RegisterSystem(sysType SystemType, sys System) error {
	return c.SystemManager.RegisterSystem(sysType, sys)
}

// SetSignature sets the signature for a system, this indicates which set of components
// the system is interested in.
func (c *Coordinator) SetSystemSignature(sysType SystemType, sig Signature) error {
	return c.SystemManager.SetSignature(sysType, sig)
}
