package entity

import (
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

// Registry keeps track of all entities and their signatures.
type Registry struct {
	nextID      uint
	signatures  [MaxEntities]Signature
	entityCount uint32
}

// NewRegistry creates a new entity registry.
func NewRegistry() *Registry {
	return &Registry{
		nextID:      0,
		signatures:  [MaxEntities]*bitset.BitSet{},
		entityCount: 0,
	}
}

func (em *Registry) Create() (Entity, error) {
	if em.entityCount >= MaxEntities {
		return 0, fmt.Errorf("too many entities")
	}

	id := Entity(em.nextID)
	em.nextID++

	em.signatures[id] = bitset.New(MaxComponents)
	em.entityCount++

	return id, nil
}

func (em *Registry) Destroy(id Entity) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id].ClearAll()
	em.entityCount--

	return nil
}

func (em *Registry) SetSignature(id Entity, sig Signature) error {
	if uint32(id) >= MaxEntities {
		return fmt.Errorf("invalid entity ID")
	}

	em.signatures[id] = sig
	return nil
}

func (em *Registry) Signature(id Entity) (Signature, error) {
	if uint32(id) >= MaxEntities {
		return nil, fmt.Errorf("invalid entity ID")
	}

	return em.signatures[id], nil
}
