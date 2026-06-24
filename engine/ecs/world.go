package ecs

import (
	"fmt"
)

// World owns all entities and a single Registry holding every
// component pool.
type World struct {
	entities      []entity
	generations   []uint32   // mirrors entities slice but tracks entity generations
	destroyQueue  []EntityID // entities to be destroyed
	freedEntities []EntityID // entityIDs made availabble after destrution
	nextIndex     uint32
	Registry      *Registry // exposed to allow direct access - may change
	Events        *EventBus
	systems       []System
	renders       []Renderer
	Logger        Logger
	Flags         map[string]bool
}

func NewWorld() *World {
	return &World{
		entities:    []entity{},
		generations: []uint32{},
		Registry:    NewRegistry(),
		Events:      NewEventBus(),
		Logger:      defaultLogger{},
		Flags:       map[string]bool{},
	}
}

func (w *World) Update(dt float64) {
	// update systems first in the event
	// of deleted entities
	for _, sys := range w.systems {
		sys.Update(w, dt)
	}
	w.Events.HandleEvents()
	// delete any entities in the queue
	for _, e := range w.destroyQueue {
		w.destroy(e)
	}
	w.destroyQueue = w.destroyQueue[:0]
}

func (w *World) destroy(e EntityID) {
	idx := e.Index()
	if int(idx) >= len(w.generations) || w.generations[idx] != e.Gen() {
		w.Logger.Warn("DestroyEntity called with stale or invalid EntityID -- ignored")
		return
	}
	w.generations[idx]++
	w.Registry.RemoveEntity(e)

	recycled := NewEntityID(idx, w.generations[idx])
	w.freedEntities = append(w.freedEntities, recycled)
}

func (w *World) IsAlive(e EntityID) bool {
	idx := e.Index()
	return int(idx) < len(w.generations) && w.generations[idx] == e.Gen()
}

func (w *World) Render() {
	for _, r := range w.renders {
		r.Render(w)
	}
}

func (w *World) RegisterSystem(s System) {
	w.systems = append(w.systems, s)
	if renderer, ok := s.(Renderer); ok {
		w.renders = append(w.renders, renderer)
	}
	if subscriber, ok := s.(EventSubscriber); ok {
		subscriber.RegisterHandlers(w)
	}
}

func (w *World) SetLogger(l Logger) {
	w.Logger = l
}

func (w *World) NewEntity() EntityID {
	if n := len(w.freedEntities); n > 0 {
		// generation was already bumped at
		// destroy time, so the ID stored in
		// freedEntities is ready to use.
		id := w.freedEntities[n-1]
		w.freedEntities = w.freedEntities[:n-1]
		w.entities[id.Index()] = entity{id: id}
		w.Logger.Info(fmt.Sprintf("entity %d created (recycled)", id))
		return id
	}
	idx := w.nextIndex
	id := NewEntityID(idx, 0)
	w.generations = append(w.generations, 0)
	w.entities = append(w.entities, entity{id: id})
	w.nextIndex++
	// w.Logger.Info(fmt.Sprintf("entity %d created", id))
	return id
}

func (w *World) DestroyEntity(e EntityID) {
	w.destroyQueue = append(w.destroyQueue, e)
}
