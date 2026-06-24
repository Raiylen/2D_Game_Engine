package systems

import (
	"github.com/raiylen/2d_game_engine/engine/components"
	"github.com/raiylen/2d_game_engine/engine/ecs"
	"github.com/veandco/go-sdl2/sdl"
)

type animationSystem struct{}

func NewAnimationSystem() *animationSystem {
	return &animationSystem{}
}

func (a *animationSystem) Update(w *ecs.World, dt float64) {
	now := int(sdl.GetTicks())
	view := ecs.NewView2[components.Animation, components.Sprite](w)
	view.Each(func(e ecs.EntityID, animation *components.Animation, sprite *components.Sprite) {
		if animation.StartTime == 0 {
			animation.StartTime = now
		}
		elapsed := now - animation.StartTime
		frameSpeed := 1000 / animation.FrameRate
		totalFrames := elapsed / frameSpeed
		if animation.Loop {
			animation.CurrentFrame = totalFrames % animation.NumFrames
		} else {
			frame := totalFrames
			if frame >= animation.NumFrames {
				frame = animation.NumFrames - 1
			}
			animation.CurrentFrame = frame
		}
		sprite.SourcePosX = animation.CurrentFrame * sprite.Width
	})
}
