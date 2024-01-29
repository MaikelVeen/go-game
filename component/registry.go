package component

import (
	"fmt"

	"github.com/MaikelVeen/go-game/entity"
)

// Registry is a struct that holds all the components types
// and their underlying stores.
type Registry struct {
	componentStores map[Type]*Store[any]
}

// NewRegistry creates a new component registry.
func NewRegistry() *Registry {
	return &Registry{
		componentStores: map[Type]*Store[any]{
			0: NewStore[any](), // Transform
			1: NewStore[any](), // SpriteRender
			2: NewStore[any](), // PlayerController
			3: NewStore[any](), // BoxCollider
			4: NewStore[any](), // Rigidbody
		},
	}
}

func (cm *Registry) AddComponent(entity entity.Entity, componentType Type, component any) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Add(entity, component)
}

func (cm *Registry) RemoveComponent(entity entity.Entity, componentType Type) error {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return fmt.Errorf("component type not registered")
	}
	return store.Remove(entity)
}

func (cm *Registry) GetComponent(entity entity.Entity, componentType Type) (any, error) {
	store, exists := cm.componentStores[componentType]
	if !exists {
		return nil, fmt.Errorf("component type not registered")
	}
	return store.Component(entity)
}

func (cm *Registry) EntityDestroyed(entity entity.Entity) {
	for _, store := range cm.componentStores {
		store.EntityDestroy(entity)
	}
}
