package systems

import (
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type damageSystem struct{}

func NewDamageSystem() *damageSystem {
	return &damageSystem{}
}

func (d *damageSystem) RegisterHandlers(w *ecs.World) {
	w.Events.RegisterHandler("collision", func(e ecs.Event) {
		w.Logger.Info("Collision Callback Functioning!")
	})
}
func (d *damageSystem) Update(w *ecs.World, dt float64) {
	//TODO: Add update method body
}
