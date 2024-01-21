package ecs

import (
	"reflect"
)

// TODO: Get Rid of the reflect package.

// SystemManager manages systems and their signatures
type SystemManager struct {
	systems    map[reflect.Type]System
	signatures map[reflect.Type]Signature
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems:    make(map[reflect.Type]System),
		signatures: make(map[reflect.Type]Signature),
	}
}

func (sm *SystemManager) Systems() []System {
	systems := make([]System, 0)
	for _, sys := range sm.systems {
		systems = append(systems, sys)
	}

	return systems
}

func (sm *SystemManager) RegisterSystem(sys System) {
	sysType := reflect.TypeOf(sys)
	if _, exists := sm.systems[sysType]; exists {
		panic("Registering system more than once")
	}

	sm.systems[sysType] = sys
}

func (sm *SystemManager) SetSignature(sys System, sig Signature) {
	sysType := reflect.TypeOf(sys)
	if _, exists := sm.systems[sysType]; !exists {
		panic("System used before registered")
	}

	sm.signatures[sysType] = sig
}

func (sm *SystemManager) EntityDestroyed(entity Entity) {
	for _, sys := range sm.systems {
		sys.EntityDestroyed(entity)
	}
}

// TODO: This could be done with events.
func (sm *SystemManager) EntitySignatureChanged(entity Entity, sig Signature) {
	for sysType, sys := range sm.systems {
		systemSignature := sm.signatures[sysType]

		if systemSignature.Intersection(sig).Count() == 0 {
			sys.EntityDestroyed(entity)
		} else {
			sys.AddEntity(entity)
		}
	}
}
