package ecs

import (
	"fmt"
	"reflect"
)

// ComponentManager is a struct that holds all the components types
// and their underlying stores.
type ComponentManager struct {
	types    map[reflect.Type]ComponentType
	stores   map[reflect.Type]Store
	nextType ComponentType
}

// NewComponentManager creates a new component registry.
func NewComponentManager() *ComponentManager {
	return &ComponentManager{
		types:    make(map[reflect.Type]ComponentType),
		stores:   make(map[reflect.Type]Store),
		nextType: 1,
	}
}

// RegisterType registers a component type and its store.
func (cm *ComponentManager) RegisterType(componentType reflect.Type, store Store) error {
	if _, exists := cm.types[componentType]; exists {
		return fmt.Errorf("component type already registered")
	}

	cm.types[componentType] = cm.nextType
	cm.stores[componentType] = store
	cm.nextType++

	return nil
}

// Type returns the component type for a given component type.
func (cm *ComponentManager) Type(componentType reflect.Type) (ComponentType, error) {
	ct, exists := cm.types[componentType]
	if !exists {
		return 0, fmt.Errorf("component type not registered")
	}

	return ct, nil
}

func (cm *ComponentManager) AddComponent(entity Entity, componentType reflect.Type, component any) error {
	store, exists := cm.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.AddComponent(entity, component)
}

func (cm *ComponentManager) RemoveComponent(entity Entity, componentType reflect.Type) error {
	store, exists := cm.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.RemoveComponent(entity)
}

func (cm *ComponentManager) GetComponent(entity Entity, componentType reflect.Type) (any, error) {
	store, exists := cm.stores[componentType]
	if !exists {
		return nil, fmt.Errorf("component type not registered")
	}
	return store.GetComponent(entity)
}

func (cm *ComponentManager) EntityDestroyed(entity Entity) {
	for _, store := range cm.stores {
		store.EntityDesroyed(entity)
	}
}
