package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type ProjectileSystem struct{}

func NewProjectileSystem() *ProjectileSystem {
	return &ProjectileSystem{}
}

func (p *ProjectileSystem) Update(w *ecs.World, dt float64) {
	view := ecs.NewView1[components.Projectile](w)
	view.Each(func(e ecs.EntityID, proj *components.Projectile) {
		proj.Timer += dt
		if proj.Timer >= proj.Duration {
			w.DestroyEntity(e)
		}
	})
}
