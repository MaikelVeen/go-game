package entity

// MaxEntities is the maximum number of entities that can be created.
const MaxEntities uint32 = 1024

// MaxComponents is the maximum number of components that can be registered per entity.
const MaxComponents uint = 32

// Entity is a unique identifier for a game object.
type Entity uint32
