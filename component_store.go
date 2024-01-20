package main

import "fmt"

type ComponentStore[T any] struct {
	// Packed slice of components set to a sp
	components [MaxEntities]T
	// Map from an entity ID to an array index.
	entityToIndexMap map[Entity]int
	// Map from an array index to an entity ID.
	indexToEntityMap map[int]Entity
	// Total size of valid entries in the array.
	size int
}

// NewComponentStore creates a new component store.
func NewComponentStore[T any]() *ComponentStore[T] {
	return &ComponentStore[T]{
		components:       [MaxEntities]T{},
		entityToIndexMap: make(map[Entity]int),
		indexToEntityMap: make(map[int]Entity),
		size:             0,
	}
}

// Add adds a component to an entity.
func (cs *ComponentStore[T]) Add(entity Entity, component T) error {
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
func (cs *ComponentStore[T]) Remove(entity Entity) error {
	indexOfRemovedEntity, exists := cs.entityToIndexMap[entity]
	if !exists {
		return fmt.Errorf("removing non-existent component")
	}

	// Copy element at end into deleted element's place to maintain density
	lastIndex := cs.size - 1
	cs.components[indexOfRemovedEntity] = cs.components[lastIndex]

	// Update map to point to moved spot
	lastEntity := cs.indexToEntityMap[lastIndex]
	cs.entityToIndexMap[lastEntity] = indexOfRemovedEntity

	// Remove map entries for deleted entity
	delete(cs.entityToIndexMap, entity)
	delete(cs.indexToEntityMap, lastIndex)

	cs.size--
	return nil
}

// Component gets the component data for an entity.
func (ca *ComponentStore[T]) Component(entity Entity) (T, error) {
	index, exists := ca.entityToIndexMap[entity]
	if !exists {
		var zero T
		return zero, fmt.Errorf("retrieving non-existent component")
	}
	return ca.components[index], nil
}
