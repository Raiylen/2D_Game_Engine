package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/raiylen/2d_game_engine/engine/events"
)

type playerSystem struct{}

func NewPlayerSystem() *playerSystem {
	return &playerSystem{}
}

func (p *playerSystem) RegisterHandlers(w *ecs.World) {
	w.Events.RegisterHandler(events.PlayerControl, func(e ecs.Event) {
		data, ok := e.Data.(events.PlayerControlEvent)
		if !ok {
			w.Logger.Warn("control event recieved unexpected data type")
			return
		}
		view := ecs.NewView3[components.Player, components.Sprite, components.RigidBody](w)
		view.Each(func(e ecs.EntityID, plr *components.Player, spr *components.Sprite, vel *components.RigidBody) {
			if data.Up {
				vel.DY = -plr.UpVelocity
				vel.DX = 0
				spr.SourcePosY = spr.Height * 0
				plr.FacingX, plr.FacingY = 0, -1
			}
			if data.Down {
				vel.DY = plr.DownVelocity
				vel.DX = 0
				spr.SourcePosY = spr.Height * 2
				plr.FacingX, plr.FacingY = 0, 1
			}
			if data.Left {
				vel.DX = -plr.LeftVelocity
				vel.DY = 0
				spr.SourcePosY = spr.Height * 3
				plr.FacingX, plr.FacingY = -1, 0
			}
			if data.Right {
				vel.DX = plr.RightVelocity
				vel.DY = 0
				spr.SourcePosY = spr.Height * 1
				plr.FacingX, plr.FacingY = 1, 0
			}
		})
	})
}

func (c *playerSystem) Update(w *ecs.World, dt float64) {}
