package main

// Signature represents a bitset of component types.
type Signature struct {
	// Bitset is a bitset of component types.
	Bitset uint64
}

// New creates a new signature.
func New(maxComponents uint8) *Signature {
	return &Signature{
		Bitset: 0,
	}
}
