package ecs

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

// EntityManager manages entities.
type EntityManager struct {
	nextID      uint
	signatures  [MaxEntities]Signature
	entityCount uint32
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		nextID:      0,
		signatures:  [MaxEntities]*bitset.BitSet{},
		entityCount: 0,
	}
}

func (em *EntityManager) Create() (Entity, error) {
	if em.entityCount >= MaxEntities {
		return 0, fmt.Errorf("too many entities")
	}

	id := Entity(em.nextID)
	em.nextID++

	em.signatures[id] = bitset.New(MaxComponents)
	em.entityCount++

	return id, nil
}

func (em *EntityManager) Destroy(id Entity) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id].ClearAll()
	em.entityCount--

	return nil
}

func (em *EntityManager) SetSignature(id Entity, sig Signature) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id] = sig
	return nil
}

func (em *EntityManager) Signature(id Entity) (Signature, error) {
	if uint32(id) >= MaxEntities {
		return nil, fmt.Errorf("invalid entity ID")
	}

	return em.signatures[id], nil
}
