package system

import (
	"errors"

	"github.com/MaikelVeen/go-game/entity"
)

var (
	ErrSystemAlreadyRegistered = errors.New("cannot register system, SystemType already registered")
	ErrSystemNotRegistered     = errors.New("system not registered")
)

// Registry manages systems and their signatures
type Registry struct {
	systems    map[Type]System
	signatures map[Type]entity.Signature
}

func NewRegistry() *Registry {
	return &Registry{
		systems:    make(map[Type]System),
		signatures: make(map[Type]entity.Signature),
	}
}

// ForEachSystem iterates over all systems and calls the provided function.
// If the function returns an error, the iteration is stopped and the error is returned.
func (sm *Registry) ForEachSystem(fn func(System) error) error {
	for _, sys := range sm.systems {
		if err := fn(sys); err != nil {
			return err
			// TODO: Think about the notion of recoverabla errors and
			// which class should be responsible for handling them.
		}
	}
	return nil
}

// RegisterSystem registers a system with the SystemManager.
// Returns an error if the system is already registered.
func (sm *Registry) RegisterSystem(sysType Type, sys System) error {
	if _, exists := sm.systems[sysType]; exists {
		return ErrSystemAlreadyRegistered
	}

	sm.systems[sysType] = sys
	return nil
}

// SetSignature sets the signature for a system, this indicates which set of components
// the system is interested in.
func (sm *Registry) SetSignature(sysType Type, sig entity.Signature) error {
	if _, exists := sm.systems[sysType]; !exists {
		return ErrSystemNotRegistered
	}

	sm.signatures[sysType] = sig
	return nil
}

// EntityDestroyed notifies all systems that an entity has been destroyed.
func (sm *Registry) EntityDestroyed(entity entity.Entity) {
	for _, sys := range sm.systems {
		sys.EntityDestroyed(entity)
	}
}

// TODO: Entiry should be added if it contains all the components the system is interested in.
// And not only when it contains an exact match of the system's signature.
// EntitySignatureChanged notifies all systems that an entity's signature has changed.
func (sm *Registry) EntitySignatureChanged(entity entity.Entity, sig entity.Signature) {
	for sysType, sys := range sm.systems {
		systemSignature := sm.signatures[sysType]

		// Check if all bits in the system's signature are set in the entity's signature.
		intersection := sig.Intersection(systemSignature)
		if intersection.Equal(systemSignature) {
			sys.AddEntity(entity)
		} else {
			sys.EntityDestroyed(entity)
		}
	}
}
