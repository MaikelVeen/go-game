package ecs

import (
	"fmt"
	"reflect"
)

// ComponetRegistery is a struct that holds all the components types
// and their stores.
type ComponentRegistry struct {
	types    map[reflect.Type]ComponentType
	stores   map[reflect.Type]Store
	nextType ComponentType
}

// NewComponentRegistry creates a new component registry.
func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		types:    make(map[reflect.Type]ComponentType),
		stores:   make(map[reflect.Type]Store),
		nextType: 1,
	}
}

// RegisterType registers a component type and its store.
func (cm *ComponentRegistry) RegisterType(componentType reflect.Type, store Store) error {
	if _, exists := cm.types[componentType]; exists {
		return fmt.Errorf("component type already registered")
	}

	cm.types[componentType] = cm.nextType
	cm.stores[componentType] = store
	cm.nextType++

	return nil
}

// Type returns the component type for a given component type.
func (cr *ComponentRegistry) Type(componentType reflect.Type) (ComponentType, error) {
	ct, exists := cr.types[componentType]
	if !exists {
		return 0, fmt.Errorf("component type not registered")
	}

	return ct, nil
}

func (cr *ComponentRegistry) AddComponent(entity Entity, componentType reflect.Type, component any) error {
	store, exists := cr.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.AddComponent(entity, component)
}

func (cr *ComponentRegistry) RemoveComponent(entity Entity, componentType reflect.Type) error {
	store, exists := cr.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.RemoveComponent(entity)
}

func (cr *ComponentRegistry) GetComponent(entity Entity, componentType reflect.Type) (any, error) {
	store, exists := cr.stores[componentType]
	if !exists {
		return nil, fmt.Errorf("component type not registered")
	}
	return store.GetComponent(entity)
}

func (cr *ComponentRegistry) EntityDestroyed(entity Entity) {
	for _, store := range cr.stores {
		store.EntityDesroyed(entity)
	}
}
