package component

import (
	"fmt"

	"github.com/MaikelVeen/go-game/entity"
)

type Store[T any] struct {
	// Packed slice of components set to a sp
	components [entity.MaxEntities]T
	// Map from an entity ID to an array index.
	entityToIndexMap map[entity.Entity]int
	// Map from an array index to an entity ID.
	indexToEntityMap map[int]entity.Entity
	// Total size of valid entries in the array.
	size int
}

// NewStore creates a new component store.
func NewStore[T any]() *Store[T] {
	return &Store[T]{
		components:       [entity.MaxEntities]T{},
		entityToIndexMap: make(map[entity.Entity]int),
		indexToEntityMap: make(map[int]entity.Entity),
		size:             0,
	}
}

// Add adds a component to an entity.
func (s *Store[T]) Add(entity entity.Entity, component T) error {
	if _, exists := s.entityToIndexMap[entity]; exists {
		return fmt.Errorf("component added to same entity more than once")
	}

	newIndex := s.size
	s.entityToIndexMap[entity] = newIndex
	s.indexToEntityMap[newIndex] = entity
	s.components[newIndex] = component
	s.size++

	return nil
}

// Remove removes a component from an entity.
func (s *Store[T]) Remove(entity entity.Entity) error {
	indexOfRemovedEntity, exists := s.entityToIndexMap[entity]
	if !exists {
		return fmt.Errorf("removing non-existent component")
	}

	indexOfLastElement := s.size - 1
	s.components[indexOfRemovedEntity] = s.components[indexOfLastElement]

	entityOfLastElement := s.indexToEntityMap[indexOfLastElement]
	s.entityToIndexMap[entityOfLastElement] = indexOfRemovedEntity
	s.indexToEntityMap[indexOfRemovedEntity] = entityOfLastElement

	delete(s.entityToIndexMap, entity)
	delete(s.indexToEntityMap, indexOfLastElement)

	s.size--
	return nil
}

// Component gets the component data for an entity.
func (ca *Store[T]) Component(entity entity.Entity) (T, error) {
	index, exists := ca.entityToIndexMap[entity]
	if !exists {
		var zero T
		return zero, fmt.Errorf("retrieving non-existent component")
	}
	return ca.components[index], nil
}

// EnityDestroy is called when an entity is destroyed.
func (s *Store[T]) EntityDestroy(entity entity.Entity) {
	if _, exists := s.entityToIndexMap[entity]; !exists {
		return
	}

	s.Remove(entity)
}
