package ecs

import "github.com/bits-and-blooms/bitset"

type Signature = *bitset.BitSet

func NewSignature(components ...ComponentType) Signature {
	signature := bitset.New(MaxComponents)
	for _, component := range components {
		signature.Set(uint(component))
	}

	return signature
}
