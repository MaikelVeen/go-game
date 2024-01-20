package ecs

import (
	"reflect"
	"sync"
)

// SystemManager manages systems and their signatures
type SystemManager struct {
	mu         sync.RWMutex // for safe concurrent access
	systems    map[reflect.Type]System
	signatures map[reflect.Type]Signature
}

func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems:    make(map[reflect.Type]System),
		signatures: make(map[reflect.Type]Signature),
	}
}

func (sm *SystemManager) RegisterSystem(sys System) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sysType := reflect.TypeOf(sys)
	if _, exists := sm.systems[sysType]; exists {
		panic("Registering system more than once")
	}

	sm.systems[sysType] = sys
}

func (sm *SystemManager) SetSignature(sys System, signature Signature) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sysType := reflect.TypeOf(sys)
	if _, exists := sm.systems[sysType]; !exists {
		panic("System used before registered")
	}

	sm.signatures[sysType] = signature
}

func (sm *SystemManager) EntityDestroyed(entity Entity) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, sys := range sm.systems {
		sys.EntityDestroyed(entity)
	}
}

func (sm *SystemManager) EntitySignatureChanged(entity Entity, entitySignature Signature) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for sysType, sys := range sm.systems {
		systemSignature := sm.signatures[sysType]

		// Entity signature matches system signature - insert into set
		// Assuming systems have a way to manage their entities
		if (uint(entitySignature) & uint(systemSignature)) == uint(systemSignature) {
			sys.AddEntity(entity)
		} else {
			sys.EntityDestroyed(entity)
		}
	}
}
