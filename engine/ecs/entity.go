package ecs

import (
// "github.com/raiylen/2d_game_engine/engine/components"
)

// EntityID is just an index -- entities have no behavior of their own.
type EntityID uint32

// entity is currently just a wrapper around an ID. It's a placeholder
// to store additional entity data that may be added.
type entity struct {
	id        EntityID
	destroyed bool
}
