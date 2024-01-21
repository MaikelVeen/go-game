package ecs

import (
	"testing"

	"github.com/bits-and-blooms/bitset"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestEntitySignatureChanged(t *testing.T) {
	t.Parallel()

	systemManager := NewSystemManager()
	mockSystem := &MockSystem{
		entities: make(map[Entity]struct{}),
	}

	assert.NoError(t, systemManager.RegisterSystem(0, mockSystem))
	// The mock system is only interested in entities with components 5 and 10.
	assert.NoError(t, systemManager.SetSignature(0, bitset.New(MaxComponents).Set(5).Set(10)))

	entity := Entity(1)
	signature := bitset.New(MaxComponents).Set(5)

	systemManager.EntitySignatureChanged(entity, signature)
	assert.Len(t, mockSystem.entities, 0)

	signature.Set(10)
	systemManager.EntitySignatureChanged(entity, signature)
	assert.Len(t, mockSystem.entities, 1)

	// If another component is added that the system is not interested in,
	// the entity should be remain because all the components the system is interested in are still present.
	signature.Set(15)
	systemManager.EntitySignatureChanged(entity, signature)
	assert.Len(t, mockSystem.entities, 1)

	// If a component is removed that the system is interested in,
	// the entity should be removed from the system.
	signature.Clear(10)
	systemManager.EntitySignatureChanged(entity, signature)
	assert.Len(t, mockSystem.entities, 0)
}

type MockSystem struct {
	entities map[Entity]struct{}
}

func (ms *MockSystem) AddEntity(entity Entity) {
	ms.entities[entity] = struct{}{}
}

func (ms *MockSystem) EntityDestroyed(entity Entity) {
	delete(ms.entities, entity)
}

func (ms *MockSystem) Update() error      { return nil }
func (ms *MockSystem) Draw(*ebiten.Image) {}
