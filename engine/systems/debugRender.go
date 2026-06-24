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
	view := ecs.NewView2[components.BoxCollider, components.Transform](w)
	view.Each(func(e ecs.EntityID, box *components.BoxCollider, pos *components.Transform) {
		srcRect := sdl.Rect{
			X: int32(pos.X + float64(box.OffsetX)),
			Y: int32(pos.Y + float64(box.OffsetY)),
			W: int32(box.Width),
			H: int32(box.Height),
		}
		d.renderer.SetDrawColor(255, 0, 0, 255)
		d.renderer.DrawRect(&srcRect)
	})
}
