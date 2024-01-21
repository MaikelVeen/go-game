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

func (c *Coordinator) RegisterSystem(sys System) {
	c.SystemManager.RegisterSystem(sys)
}

func (c *Coordinator) SetSystemSignature(sys System, sig Signature) {
	c.SystemManager.SetSignature(sys, sig)
}

func (c *Coordinator) Update() error {
	for _, sys := range c.SystemManager.Systems() {
		if err := sys.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Coordinator) Draw(screen *ebiten.Image) {
	for _, sys := range c.SystemManager.Systems() {
		sys.Draw(screen)
	}
}
