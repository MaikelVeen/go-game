package ecs

// Signature represents a bitset of component types.
type Signature uint64

// NewSignature creates a new signature.
func NewSignature(componentTypes ...ComponentType) Signature {
	var signature Signature
	for _, componentType := range componentTypes {
		signature.Set(componentType)
	}
	return signature
}

// Set adds a component type to the signature.
func (s *Signature) Set(componentType ComponentType) {
	*s |= 1 << componentType
}
