package ecs

import (
// "github.com/raiylen/2d_game_engine/engine/components"
)

// EntityID is just an index -- entities have no behavior of their own.
// The slot index (lower 32 bits) and generation counter (upper 32 bits)
// are rolled into the same 64 bit numbeer. Two IDs referring to the same
// slot but different generations are not equal
type EntityID uint64

func NewEntityID(index, gen uint32) EntityID {
	return EntityID(uint64(gen)<<32 | uint64(index))
}

// Index returns the entityID number
func (e EntityID) Index() uint32 {
	return uint32(e)
}

// Gen returns the generation number -- used to detect stale handles
func (e EntityID) Gen() uint32 {
	return uint32(e >> 32)
}

// entity is currently just a wrapper around an ID. It's a placeholder
// to store additional entity data that may be added.
type entity struct {
	id EntityID
}
