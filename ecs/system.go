package ecs

type System interface {
	AddEntity(entity Entity)
	EntityDestroyed(entity Entity)
	Update()
}
