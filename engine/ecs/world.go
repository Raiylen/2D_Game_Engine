package ecs

import (
// "fmt"
)

// World owns all entities and a single Registry holding every
// component pool.
type World struct {
	entities      []entity
	destroyQueue  []EntityID // entities to be destroyed
	freedEntities []EntityID // entityIDs made availabble after destrution
	nextID        EntityID
	Registry      *Registry // exposed to allow direct access - may change
	systems       []System
	renders       []Renderer
	Logger        Logger
	Flags         map[string]bool
}

func NewWorld() *World {
	return &World{
		entities: []entity{},
		Registry: NewRegistry(),
		Logger:   defaultLogger{},
		Flags:    map[string]bool{},
	}
}

func (w *World) Update(dt float64) {
	// update systems first in the event
	// of deleted entities
	for _, sys := range w.systems {
		sys.Update(w, dt)
	}
	// delete any entities in the queue
	for _, e := range w.destroyQueue {
		if w.entities[e].destroyed {
			continue // already removed -- avoids double-freeing the ID
		}
		w.entities[e].destroyed = true
		w.Registry.RemoveEntity(e)
		w.freedEntities = append(w.freedEntities, e)
	}
	w.destroyQueue = w.destroyQueue[:0]
}

func (w *World) Render() {
	for _, r := range w.renders {
		r.Render(w)
	}
}

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
}

func (w *World) RegisterRender(r Renderer) {
	w.renders = append(w.renders, r)
}

func (w *World) SetLogger(l Logger) {
	w.Logger = l
}

func (w *World) NewEntity() EntityID {
	if n := len(w.freedEntities); n > 0 {
		id := w.freedEntities[n-1]
		w.freedEntities = w.freedEntities[:n-1]
		w.entities[id] = entity{id: id}
		// w.Logger.Info(fmt.Sprintf("entity %d created (recycled)", id))
		return id
	}
	id := w.nextID
	w.entities = append(w.entities, entity{id: id})
	w.nextID++
	// w.Logger.Info(fmt.Sprintf("entity %d created", id))
	return id
}

func (w *World) DestroyEntity(e EntityID) {
	w.destroyQueue = append(w.destroyQueue, e)
}
