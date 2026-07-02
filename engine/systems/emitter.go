package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/raiylen/2d_game_engine/engine/events"
)

type EmitterSystem struct{}

func NewEmitterSystem() *EmitterSystem {
	return &EmitterSystem{}
}

func (em *EmitterSystem) RegisterHandlers(w *ecs.World) {
	w.Events.RegisterHandler(events.KbFire, func(e ecs.Event) {
		_, ok := e.Data.(events.KbFireEvent)
		if !ok {
			w.Logger.Warn("fire event recieved unexpected data type")
			return
		}
		view := ecs.NewView4[components.Player, components.Emitter, components.Transform, components.Sprite](w)
		view.Each(func(e ecs.EntityID, player *components.Player, emit *components.Emitter, pos *components.Transform, spr *components.Sprite) {
			emit.DX = emit.ProjSpeed * player.FacingX
			emit.DY = emit.ProjSpeed * player.FacingY

			em.createProjectile(w, emit, pos, spr)
		})
	})
}
func (em *EmitterSystem) Update(w *ecs.World, dt float64) {
	view := ecs.NewView3[components.Emitter, components.Transform, components.Sprite](w)
	view.Each(func(e ecs.EntityID, emit *components.Emitter, pos *components.Transform, spr *components.Sprite) {
		if emit.Frequency == 0 {
			return
		}
		emit.Timer += dt
		secondsPerShot := 1 / emit.Frequency
		if emit.Timer >= secondsPerShot {
			emit.Timer = 0
			em.createProjectile(w, emit, pos, spr)
		}
	})
}

func (em *EmitterSystem) createProjectile(w *ecs.World, emit *components.Emitter, pos *components.Transform, spr *components.Sprite) {
	projSpr := components.Sprite{AssetID: "bullet-image", Width: 4, Height: 4, Layer: 11}

	srcCtrX := pos.X + float64(spr.Width/2-projSpr.Width/2)
	srcCtrY := pos.Y + float64(spr.Height/2-projSpr.Height/2)

	projectile := w.NewEntity()
	ecs.AddComponent(w.Registry, projectile, components.Transform{X: srcCtrX, Y: srcCtrY})
	ecs.AddComponent(w.Registry, projectile, components.RigidBody{DX: emit.DX, DY: emit.DY})
	ecs.AddComponent(w.Registry, projectile, components.Projectile{
		Duration:   emit.ProjDuration,
		PercDamage: emit.ProjPercDamage,
		IsFriendly: emit.IsFriendly,
	})
	ecs.AddComponent(w.Registry, projectile, projSpr)
}
