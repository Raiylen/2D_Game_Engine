package systems

import (
	"fmt"
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type colliderEntry struct {
	entity ecs.EntityID
	box    *components.BoxCollider
	pos    *components.Transform
}

type collisionSystem struct {
	colliders []colliderEntry
}

func NewCollisionSystem() *collisionSystem {
	return &collisionSystem{}
}

func (c *collisionSystem) Update(w *ecs.World, dt float64) {
	view := ecs.NewView2[components.BoxCollider, components.Transform](w)
	c.colliders = c.colliders[:0]
	view.Each(func(e ecs.EntityID, box *components.BoxCollider, pos *components.Transform) {
		c.colliders = append(c.colliders, colliderEntry{entity: e, box: box, pos: pos})
	})

	for i := 0; i < len(c.colliders); i++ {
		for j := i + 1; j < len(c.colliders); j++ {
			a := c.colliders[i]
			b := c.colliders[j]

			if c.checkCollision(a, b) {
				w.Logger.Info(fmt.Sprintf("Collision detected! %v, %v", a.entity, b.entity))
			}
		}
	}
}

func (c *collisionSystem) checkCollision(a, b colliderEntry) bool {
	aMinX := a.pos.X + float64(a.box.OffsetX)
	aMaxX := a.pos.X + float64(a.box.Width+a.box.OffsetX)
	aMinY := a.pos.Y + float64(a.box.OffsetY)
	aMaxY := a.pos.Y + float64(a.box.Height+a.box.OffsetY)

	bMinX := b.pos.X + float64(b.box.OffsetX)
	bMaxX := b.pos.X + float64(b.box.Width+b.box.OffsetX)
	bMinY := b.pos.Y + float64(b.box.OffsetY)
	bMaxY := b.pos.Y + float64(b.box.Height+b.box.OffsetY)

	return (aMinX <= bMaxX &&
		aMaxX >= bMinX &&
		aMinY <= bMaxY &&
		aMaxY >= bMinY)
}
