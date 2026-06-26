package systems

import (
	"fmt"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/raiylen/2d_game_engine/engine/events"
)

type damageSystem struct{}

func NewDamageSystem() *damageSystem {
	return &damageSystem{}
}

func (d *damageSystem) RegisterHandlers(w *ecs.World) {
	w.Events.RegisterHandler(events.Collision, func(e ecs.Event) {
		data, ok := e.Data.(events.CollisionEvent)
		if !ok {
			w.Logger.Warn("collision event recieved unexpected data type")
			return
		}
		w.Logger.Info(fmt.Sprintf("Collision detected between EntityID %d and %d!", data.A, data.B))
		w.DestroyEntity(data.A)
		w.DestroyEntity(data.B)
	})
}
func (d *damageSystem) Update(w *ecs.World, dt float64) {
	//TODO: Add update method body
}
