package systems

import (
	// "fmt"
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	// log "github.com/raiylen/2d_game_engine/logger"
)

type debugSystem struct{}

func NewDebugSystem() *debugSystem {
	return &debugSystem{}
}

func (d *debugSystem) Update(w *ecs.World, dt float32) {
	view := ecs.NewView1[components.Transform](w)
	view.Each(func(e ecs.EntityID, pos *components.Transform) {
		// log.Info(fmt.Sprintf("entity %d -> Position{%.2f, %.2f}", e, pos.X, pos.Y))
		// log.Info(fmt.Sprintf("%.2f", dt))
	})

}
