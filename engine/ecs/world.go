package ecs

import (
// "github.com/raiylen/2d_game_engine/engine/components"
)

type World struct {
	entities []entity
	nextID   EntityID
	registry *Registry
	// positions *typedPool[components.Position]
	// velocitys *typedPool[components.Velocity]
	systems []System
}

func NewWorld() *World {
	return &World{
		entities: []entity{},
		registry: NewRegistry(),
		// positions: NewTypedPool[components.Position](),
		// velocitys: NewTypedPool[components.Velocity](),
	}
}

func (w *World) Update(dt float64) {
	for _, sys := range w.systems {
		sys.Update(w, dt)
	}
}

// func (w *World) ViewPositionVelocity() View2[components.Position, components.Velocity] {
// 	return NewView2(w.positions, w.velocitys)
// }

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) NewEntity() EntityID {
	id := w.nextID
	w.entities = append(w.entities, entity{id: id})
	w.nextID++
	return id
}
