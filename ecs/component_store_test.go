package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComponentStore(t *testing.T) {
	t.Parallel()

	type TestComponent struct {
		value rune
	}

	cs := NewComponentStore[TestComponent]()

	// Add component A to entity 0
	entity0 := Entity(0)
	componentA := TestComponent{value: 'A'}
	assert.NoError(t, cs.Add(entity0, componentA))

	// Add component B to entity 1
	entity1 := Entity(1)
	componentB := TestComponent{value: 'B'}
	assert.NoError(t, cs.Add(entity1, componentB))

	// Add component C to entity 2
	entity2 := Entity(2)
	componentC := TestComponent{value: 'C'}
	assert.NoError(t, cs.Add(entity2, componentC))

	// Add component D to entity 3
	entity3 := Entity(3)
	componentD := TestComponent{value: 'D'}
	assert.NoError(t, cs.Add(entity3, componentD))

	// Remove component B from entity 1
	assert.NoError(t, cs.Remove(entity1))

	// Assert that to keep it packed, we move the last element D
	// into the spot occupied by B, and update the maps.
	assert.Equal(t, cs.components[1], componentD)
	assert.Equal(t, cs.entityToIndexMap[entity3], 1)
	assert.Equal(t, cs.indexToEntityMap[1], entity3)

	// Remove component D from entity 3
	assert.NoError(t, cs.Remove(entity3))

	// Assert that to keep it packed, we move the last element C
	// into the spot occupied by D, and update the maps.
	assert.Equal(t, cs.components[1], componentC)
	assert.Equal(t, cs.entityToIndexMap[entity2], 1)
	assert.Equal(t, cs.indexToEntityMap[1], entity2)

	// Add new component E to entity 4
	entity4 := Entity(4)
	componentE := TestComponent{value: 'E'}
	assert.NoError(t, cs.Add(entity4, componentE))

	// Assert that E was added to the end of the array
	assert.Equal(t, cs.components[2], componentE)
	assert.Equal(t, cs.entityToIndexMap[entity4], 2)
	assert.Equal(t, cs.indexToEntityMap[2], entity4)
}
