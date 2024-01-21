package ecs

import (
	"fmt"
)

// ComponentManager is a struct that holds all the components types
// and their underlying stores.
type ComponentManager struct {
	componentStores map[ComponentType]*ComponentStore[any]
}

// NewComponentManager creates a new component registry.
func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		componentStores: map[ComponentType]*ComponentStore[any]{
			0: NewComponentStore[any](), // Transform
			1: NewComponentStore[any](), // SpriteRender
			2: NewComponentStore[any](), // PlayerController
			3: NewComponentStore[any](), // BoxCollider
			4: NewComponentStore[any](), // Rigidbody
		},
	}
}

func (cm *ComponentManager) AddComponent(entity Entity, componentType ComponentType, component any) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Add(entity, component)
}

func (cm *ComponentManager) RemoveComponent(entity Entity, componentType ComponentType) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Remove(entity)
}

func (cm *ComponentManager) GetComponent(entity Entity, componentType ComponentType) (any, error) {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return nil, fmt.Errorf("component type not registered")
	}
	return store.Component(entity)
}

func (cm *ComponentManager) EntityDestroyed(entity Entity) {
	for _, store := range cm.componentStores {
		store.EntityDestroy(entity)
	}
}
