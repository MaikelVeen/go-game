package entity

import "github.com/bits-and-blooms/bitset"

type Signature = *bitset.BitSet

// NewSignature creates a new Signature from the provided component types.
// It initializes a bitset with a size of MaxComponents, and sets the bits
// corresponding to the provided component types. Each component type is
// represented by a unique bit in the bitset.
func NewSignature(componentTypes ...uint) Signature {
	signature := bitset.New(MaxComponents)
	for _, component := range componentTypes {
		signature.Set(component)
	}

	return signature
}
