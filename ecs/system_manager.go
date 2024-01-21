package ecs

import (
	"errors"
)

var (
	ErrSystemAlreadyRegistered = errors.New("System already registered")
	ErrSystemNotRegistered     = errors.New("System not registered")
)

// SystemManager manages systems and their signatures
type SystemManager struct {
	systems    map[SystemType]System
	signatures map[SystemType]Signature
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems:    make(map[SystemType]System),
		signatures: make(map[SystemType]Signature),
	}
}

// ForEachSystem iterates over all systems and calls the provided function.
// If the function returns an error, the iteration is stopped and the error is returned.
func (sm *SystemManager) ForEachSystem(fn func(System) error) error {
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
func (sm *SystemManager) RegisterSystem(sysType SystemType, sys System) error {
	if _, exists := sm.systems[sysType]; exists {
		return ErrSystemAlreadyRegistered
	}

	sm.systems[sysType] = sys
	return nil
}

// SetSignature sets the signature for a system, this indicates which set of components
// the system is interested in.
func (sm *SystemManager) SetSignature(sysType SystemType, sig Signature) error {
	if _, exists := sm.systems[sysType]; !exists {
		return ErrSystemNotRegistered
	}

	sm.signatures[sysType] = sig
	return nil
}

// EntityDestroyed notifies all systems that an entity has been destroyed.
func (sm *SystemManager) EntityDestroyed(entity Entity) {
	for _, sys := range sm.systems {
		sys.EntityDestroyed(entity)
	}
}

// EntitySignatureChanged notifies all systems that an entity's signature has changed.
func (sm *SystemManager) EntitySignatureChanged(entity Entity, sig Signature) {
	for sysType, sys := range sm.systems {
		systemSignature := sm.signatures[sysType]

		if sig.Equal(systemSignature) {
			sys.AddEntity(entity)
		} else {
			sys.EntityDestroyed(entity)
		}
	}
}
