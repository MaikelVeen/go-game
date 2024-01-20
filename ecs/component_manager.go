package ecs

import (
	"fmt"
	"reflect"
)

// ComponentManager is a struct that holds all the components types
// and their underlying stores.
type ComponentManager struct {
	componentTypes  map[reflect.Type]ComponentType
	componentStores map[reflect.Type]Store
	nextType        ComponentType
}

// NewComponentManager creates a new component registry.
func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		componentTypes:  make(map[reflect.Type]ComponentType),
		componentStores: make(map[reflect.Type]Store),
		nextType:        1,
	}
}

// RegisterType registers a component type and its store.
// Returns the component type and an error if the component type is already registered.
func (cm *ComponentManager) RegisterComponentType(componentType reflect.Type) (ComponentType, error) {
	if _, exists := cm.componentTypes[componentType]; exists {
		return 0, fmt.Errorf("component type already registered")
	}

	cm.componentTypes[componentType] = cm.nextType
	cm.componentStores[componentType] = NewComponentStore()
	cm.nextType++

	return cm.nextType - 1, nil
}

// Type returns the component type for a given component type.
func (cm *ComponentManager) Type(componentType reflect.Type) (ComponentType, error) {
	ct, exists := cm.componentTypes[componentType]
	if !exists {
		return 0, fmt.Errorf("component type not registered")
	}

	return ct, nil
}

func (cm *ComponentManager) AddComponent(entity Entity, componentType reflect.Type, component any) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Add(entity, component)
}

func (cm *ComponentManager) RemoveComponent(entity Entity, componentType reflect.Type) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Remove(entity)
}

func (cm *ComponentManager) GetComponent(entity Entity, componentType reflect.Type) (any, error) {
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
