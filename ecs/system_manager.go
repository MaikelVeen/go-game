package ecs

import (
	"reflect"
)

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

func (sm *SystemManager) SetSignature(sys System, signature Signature) {
	sysType := reflect.TypeOf(sys)
	if _, exists := sm.systems[sysType]; !exists {
		panic("System used before registered")
	}

	sm.signatures[sysType] = signature
}

func (sm *SystemManager) EntityDestroyed(entity Entity) {
	for _, sys := range sm.systems {
		sys.EntityDestroyed(entity)
	}
}

func (sm *SystemManager) EntitySignatureChanged(entity Entity, entitySignature Signature) {
	for sysType, sys := range sm.systems {
		systemSignature := sm.signatures[sysType]

		if (uint(entitySignature) & uint(systemSignature)) == uint(systemSignature) {
			sys.AddEntity(entity)
		} else {
			sys.EntityDestroyed(entity)
		}
	}
}
