package ecs

import "reflect"

// ComponetRegistery is a struct that holds all the components types
// and their stores.
type ComponentRegistry struct {
	types    map[reflect.Type]ComponentType
	stores   map[reflect.Type]Store
	nextType ComponentType
}
