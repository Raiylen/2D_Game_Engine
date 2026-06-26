package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/raiylen/2d_game_engine/engine/events"
)

type keyboardControlSystem struct{}

func NewControlSystem() *keyboardControlSystem {
	return &keyboardControlSystem{}
}

func (c *keyboardControlSystem) RegisterHandlers(w *ecs.World) {
	w.Events.RegisterHandler(events.KbControl, func(e ecs.Event) {
		data, ok := e.Data.(events.KbControlEvent)
		if !ok {
			w.Logger.Warn("control event recieved unexpected data type")
			return
		}
		view := ecs.NewView3[components.KeyboardControl, components.Sprite, components.RigidBody](w)
		view.Each(func(e ecs.EntityID, ctr *components.KeyboardControl, spr *components.Sprite, vel *components.RigidBody) {
			if data.Up {
				vel.DY = -float64(ctr.UpVelocity)
				vel.DX = 0
				spr.SourcePosY = spr.Height * 0
			}
			if data.Down {
				vel.DY = float64(ctr.DownVelocity)
				vel.DX = 0
				spr.SourcePosY = spr.Height * 2
			}
			if data.Left {
				vel.DX = -float64(ctr.LeftVelocity)
				vel.DY = 0
				spr.SourcePosY = spr.Height * 3
			}
			if data.Right {
				vel.DX = float64(ctr.RightVelocity)
				vel.DY = 0
				spr.SourcePosY = spr.Height * 1
			}
		})
	})
}

func (c *keyboardControlSystem) Update(w *ecs.World, dt float64) {}
