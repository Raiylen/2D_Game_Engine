package game

import (
	ecs "github.com/raiylen/2d_game_engine/engine/ecs"
	events "github.com/raiylen/2d_game_engine/engine/events"
	sdl "github.com/veandco/go-sdl2/sdl"
)

func (g *Game) getInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch e := event.(type) {
		case *sdl.QuitEvent:
			g.isRunning = false
		case *sdl.KeyboardEvent:
			if e.Type == sdl.KEYDOWN {
				switch e.Keysym.Scancode {
				case sdl.SCANCODE_ESCAPE:
					g.isRunning = false
				case sdl.SCANCODE_F4:
					g.isFullscreen = !g.isFullscreen
					if g.isFullscreen {
						g.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
					} else {
						g.window.SetFullscreen(0)
					}
				case sdl.SCANCODE_F3:
					g.world.Flags["debug"] = !g.world.Flags["debug"]
				case sdl.SCANCODE_SPACE:
					g.world.Events.Emit(ecs.Event{
						Name: events.KbFire,
						Data: events.KbFireEvent{},
					})
				}
			}
		}
	}
	g.world.Events.Emit(ecs.Event{
		Name: events.PlayerControl,
		Data: events.PlayerControlEvent{
			Up:    g.keystate[sdl.SCANCODE_UP] == 1,
			Down:  g.keystate[sdl.SCANCODE_DOWN] == 1,
			Left:  g.keystate[sdl.SCANCODE_LEFT] == 1,
			Right: g.keystate[sdl.SCANCODE_RIGHT] == 1,
		},
	})
}
