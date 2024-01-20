package ecs

import "fmt"

type Store interface {
	EntityDestroy(entity Entity)

	Add(entity Entity, component any) error
	Remove(entity Entity) error
	Component(entity Entity) (any, error)
}

var _ Store = &ComponentStore{}

type ComponentStore struct {
	// Packed slice of components set to a sp
	components [MaxEntities]any
	// Map from an entity ID to an array index.
	entityToIndexMap map[Entity]int
	// Map from an array index to an entity ID.
	indexToEntityMap map[int]Entity
	// Total size of valid entries in the array.
	size int
}

// NewComponentStore creates a new component store.
func NewComponentStore() *ComponentStore {
	return &ComponentStore{
		components:       [MaxEntities]any{},
		entityToIndexMap: make(map[Entity]int),
		indexToEntityMap: make(map[int]Entity),
		size:             0,
	}
}

// Add adds a component to an entity.
func (cs *ComponentStore) Add(entity Entity, component any) error {
	if _, exists := cs.entityToIndexMap[entity]; exists {
		return fmt.Errorf("component added to same entity more than once")
	}

	newIndex := cs.size
	cs.entityToIndexMap[entity] = newIndex
	cs.indexToEntityMap[newIndex] = entity
	cs.components[newIndex] = component
	cs.size++

	return nil
}

// Remove removes a component from an entity.
func (cs *ComponentStore) Remove(entity Entity) error {
	indexOfRemovedEntity, exists := cs.entityToIndexMap[entity]
	if !exists {
		return fmt.Errorf("removing non-existent component")
	}

	indexOfLastElement := cs.size - 1
	cs.components[indexOfRemovedEntity] = cs.components[indexOfLastElement]

	entityOfLastElement := cs.indexToEntityMap[indexOfLastElement]
	cs.entityToIndexMap[entityOfLastElement] = indexOfRemovedEntity
	cs.indexToEntityMap[indexOfRemovedEntity] = entityOfLastElement

	delete(cs.entityToIndexMap, entity)
	delete(cs.indexToEntityMap, indexOfLastElement)

	cs.size--
	return nil
}

// Component gets the component data for an entity.
func (ca *ComponentStore) Component(entity Entity) (any, error) {
	index, exists := ca.entityToIndexMap[entity]
	if !exists {
		var zero any
		return zero, fmt.Errorf("retrieving non-existent component")
	}
	return ca.components[index], nil
}

// EnityDestroy is called when an entity is destroyed.
func (cs *ComponentStore) EntityDestroy(entity Entity) {
	if _, exists := cs.entityToIndexMap[entity]; !exists {
		return
	}

	cs.Remove(entity)
}
