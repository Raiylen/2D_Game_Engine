package ecs

type System interface {
	Update(w *World, dt float64)
}
