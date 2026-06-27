package systems

import (
	"fmt"
	"github.com/raiylen/2d_game_engine/assets"
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/veandco/go-sdl2/sdl"
	"sort"
)

type renderEntry struct {
	entity    ecs.EntityID
	transform *components.Transform
	sprite    *components.Sprite
}

type renderSystem struct {
	renderer    *sdl.Renderer
	assets      *assets.AssetManager
	renderQueue []renderEntry
}

func NewRenderSystem(renderer *sdl.Renderer, assets *assets.AssetManager) *renderSystem {
	return &renderSystem{
		renderer: renderer,
		assets:   assets,
	}
}

func (r *renderSystem) Update(w *ecs.World, dt float64) {} // no-op to to satisfy system interface

func (r *renderSystem) Render(w *ecs.World) {

	r.buildRenderQueue(w)

	camView := ecs.NewView1[components.Camera](w)
	if camView.Len() > 1 {
		w.Logger.Warn("more than one camera entity detected -- undefined behavior")
		return
	}
	camView.Each(func(e ecs.EntityID, cam *components.Camera) { // for now there is only one camera
		r.drawQueue(w, cam)
	})
}

func (r *renderSystem) buildRenderQueue(w *ecs.World) {
	view := ecs.NewView2[components.Transform, components.Sprite](w)

	r.renderQueue = r.renderQueue[:0]
	view.Each(func(e ecs.EntityID, trans *components.Transform, spr *components.Sprite) {
		if spr.AssetID == "" {
			return
		}
		if trans.Scale == 0 {
			trans.Scale = 1
		}
		r.renderQueue = append(r.renderQueue, renderEntry{e, trans, spr})
	})

	sort.Slice(r.renderQueue, func(i, j int) bool {
		return r.renderQueue[i].sprite.Layer < r.renderQueue[j].sprite.Layer
	})
}

func (r *renderSystem) drawQueue(w *ecs.World, cam *components.Camera) {
	for _, ent := range r.renderQueue {
		sprite, err := r.assets.GetTexture(ent.sprite.AssetID)
		if err != nil {
			w.Logger.Warn(fmt.Sprintf("EntityID: %d -- %s", ent.entity, err))
			ent.sprite.AssetID = ""
		}
		// don't move fixed entities with the camera
		var camX, camY float64
		if ent.sprite.IsFixed {
			camX, camY = 0, 0
		} else {
			camX, camY = cam.X, cam.Y
		}

		srcRect := sdl.Rect{
			X: int32(ent.sprite.SourcePosX),
			Y: int32(ent.sprite.SourcePosY),
			W: int32(ent.sprite.Width),
			H: int32(ent.sprite.Height),
		}

		dstRect := sdl.Rect{
			X: int32(ent.transform.X - camX),
			Y: int32(ent.transform.Y - camY),
			W: int32(float64(ent.sprite.Width) * ent.transform.Scale),
			H: int32(float64(ent.sprite.Height) * ent.transform.Scale),
		}
		r.renderer.CopyEx(sprite, &srcRect, &dstRect, ent.transform.Rotation, nil, sdl.FLIP_NONE)
	}
}
