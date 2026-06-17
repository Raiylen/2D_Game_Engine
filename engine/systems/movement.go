package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type movementSystem struct{}

func NewMovementSystem() *movementSystem {
	return &movementSystem{}
}

func (m *movementSystem) Update(w *ecs.World, dt float64) {
	view := w.ViewPositionVelocity()
	view.Each(func(e ecs.EntityID, pos *components.Position, vel *components.Velocity) {
		pos.X = pos.X + vel.DX
		pos.Y = pos.Y + vel.DY
	})
}
