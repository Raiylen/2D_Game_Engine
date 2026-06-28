package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type projectileSystem struct{}

func NewProjectileSystem() *projectileSystem {
	return &projectileSystem{}
}

func (p *projectileSystem) Update(w *ecs.World, dt float64) {
	view := ecs.NewView3[components.Projectile, components.Transform, components.Sprite](w)
	view.Each(func(e ecs.EntityID, proj *components.Projectile, pos *components.Transform, spr *components.Sprite) {
		proj.Timer += dt
		secondsPerShot := 1 / proj.Frequency
		if proj.Timer >= secondsPerShot {
			proj.Timer = 0
			p.createProjectile(w, proj, pos, spr)
		}
	})
}

func (p *projectileSystem) createProjectile(w *ecs.World, proj *components.Projectile, pos *components.Transform, spr *components.Sprite) {
	projSpr := components.Sprite{AssetID: "bullet-image", Width: 4, Height: 4, Layer: 11}

	srcCtrX := pos.X + float64(spr.Width/2-projSpr.Width/2)
	srcCtrY := pos.Y + float64(spr.Height/2-projSpr.Height/2)

	projectile := w.NewEntity()
	ecs.AddComponent(w.Registry, projectile, components.Transform{X: srcCtrX, Y: srcCtrY})
	ecs.AddComponent(w.Registry, projectile, components.RigidBody{DX: proj.DX, DY: proj.DY})
	ecs.AddComponent(w.Registry, projectile, projSpr)
}
