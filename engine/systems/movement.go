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
	view := ecs.NewView2[components.Transform, components.RigidBody](w)
	// view := ecs.NewGroupView2[components.Position, components.Velocity](w)
	view.Each(func(e ecs.EntityID, pos *components.Transform, vel *components.RigidBody) {
		pos.X = pos.X + vel.DX*dt
		pos.Y = pos.Y + vel.DY*dt
	})
}
