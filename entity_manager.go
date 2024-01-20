package main

import (
	"fmt"
)

// EntityManager manages entities.
type EntityManager struct {
	nextID      uint
	signatures  [MaxEntities]Signature
	entityCount uint32
}

// NewEntityManager creates a new entity manager.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextID:      0,
		signatures:  [MaxEntities]Signature{},
		entityCount: 0,
	}
}

// Create creates a new entity.
func (em *EntityManager) Create() (Entity, error) {
	if em.entityCount >= MaxEntities {
		return 0, fmt.Errorf("too many entities")
	}

	id := Entity(em.nextID)
	em.nextID++

	return id, nil
}

// Destroy destroys an entity.
func (em *EntityManager) Destroy(id Entity) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id].Bitset = 0
	em.entityCount--

	return nil
}

// SetSignature sets the signature of an entity.
func (em *EntityManager) SetSignature(id Entity, signature *Signature) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id] = *signature
	return nil
}

// Signature gets the signature of an entity.
func (em *EntityManager) Signature(id Entity) (*Signature, error) {
	if uint32(id) >= MaxEntities {
		return nil, fmt.Errorf("invalid entity ID")
	}

	return &em.signatures[id], nil
}
