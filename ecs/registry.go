package ecs

import (
	"fmt"
	"reflect"
)

// ComponetRegistery is a struct that holds all the components types
// and their underlying stores.
type Registry struct {
	types    map[reflect.Type]ComponentType
	stores   map[reflect.Type]Store
	nextType ComponentType
}

// NewComponentRegistry creates a new component registry.
func NewComponentRegistry() *Registry {
	return &Registry{
		types:    make(map[reflect.Type]ComponentType),
		stores:   make(map[reflect.Type]Store),
		nextType: 1,
	}
}

// RegisterType registers a component type and its store.
func (r *Registry) RegisterType(componentType reflect.Type, store Store) error {
	if _, exists := r.types[componentType]; exists {
		return fmt.Errorf("component type already registered")
	}

	r.types[componentType] = r.nextType
	r.stores[componentType] = store
	r.nextType++

	return nil
}

// Type returns the component type for a given component type.
func (r *Registry) Type(componentType reflect.Type) (ComponentType, error) {
	ct, exists := r.types[componentType]
	if !exists {
		return 0, fmt.Errorf("component type not registered")
	}

	return ct, nil
}

func (r *Registry) AddComponent(entity Entity, componentType reflect.Type, component any) error {
	store, exists := r.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.AddComponent(entity, component)
}

func (r *Registry) RemoveComponent(entity Entity, componentType reflect.Type) error {
	store, exists := r.stores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.RemoveComponent(entity)
}

func (cr *Registry) GetComponent(entity Entity, componentType reflect.Type) (any, error) {
	store, exists := cr.stores[componentType]
	if !exists {
		return nil, fmt.Errorf("component type not registered")
	}
	return store.GetComponent(entity)
}

func (cr *Registry) EntityDestroyed(entity Entity) {
	for _, store := range cr.stores {
		store.EntityDesroyed(entity)
	}
}
