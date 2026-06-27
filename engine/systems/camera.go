package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
)

type cameraSystem struct{}

func NewCameraSystem() *cameraSystem {
	return &cameraSystem{}
}

func (c *cameraSystem) Update(w *ecs.World, dt float64) {
	view := ecs.NewView2[components.Camera, components.Transform](w)
	view.Each(func(e ecs.EntityID, cam *components.Camera, pos *components.Transform) {
		// center on the player
		cam.X = pos.X - float64(cam.Width)/2
		cam.Y = pos.Y - float64(cam.Height)/2

		// clamp left
		if cam.X < 0 {
			cam.X = 0
		}
		// clamp right
		if (cam.X + float64(cam.Width)) > float64(cam.MapWidth) {
			cam.X = float64(cam.MapWidth - cam.Width)
		}
		// clamp top
		if cam.Y < 0 {
			cam.Y = 0
		}
		// clamp bottom
		if (cam.Y + float64(cam.Height)) > float64(cam.MapHeight) {
			cam.Y = float64(cam.MapHeight - cam.Height)
		}
	})
}
