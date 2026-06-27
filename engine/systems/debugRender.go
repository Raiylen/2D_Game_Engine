package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type debugRenderSystem struct {
	renderer *sdl.Renderer
}

func NewDebugRenderSystem(renderer *sdl.Renderer) *debugRenderSystem {
	return &debugRenderSystem{
		renderer: renderer,
	}
}

func (d *debugRenderSystem) Update(w *ecs.World, dt float64) {}

func (d *debugRenderSystem) Render(w *ecs.World) {
	if !w.Flags["debug"] {
		return
	}
	camView := ecs.NewView1[components.Camera](w)
	if camView.Len() > 1 {
		w.Logger.Warn("more than one camera entity detected -- undefined behavior")
		return
	}
	var camX, camY float64
	camView.Each(func(e ecs.EntityID, cam *components.Camera) {
		camX = cam.X
		camY = cam.Y
	})

	view := ecs.NewView2[components.BoxCollider, components.Transform](w)
	view.Each(func(e ecs.EntityID, box *components.BoxCollider, pos *components.Transform) {
		srcRect := sdl.Rect{
			X: int32(pos.X + float64(box.OffsetX) - camX),
			Y: int32(pos.Y + float64(box.OffsetY) - camY),
			W: int32(float64(box.Width) * pos.Scale),
			H: int32(float64(box.Height) * pos.Scale),
		}
		d.renderer.SetDrawColor(255, 0, 0, 255)
		d.renderer.DrawRect(&srcRect)
	})
}
