package game

import (
	"fmt"
	img "github.com/veandco/go-sdl2/img"
	sdl "github.com/veandco/go-sdl2/sdl"
)

func (g *Game) Init() error {
	var err error
	var sdlFlags uint32 = sdl.INIT_EVERYTHING
	var imgFlags int = img.INIT_PNG

	if err = sdl.Init(sdlFlags); err != nil {
		return fmt.Errorf("Error initializing SDL2: %v", err)
	}

	if err = img.Init(imgFlags); err != nil {
		return fmt.Errorf("Error initializing SDL_IMG: %v", err)
	}

	// displayMode, err := sdl.GetCurrentDisplayMode(0)
	// if err != nil {
	// 	return fmt.Errorf("Error retrieving current display mode: %v", err)
	// }

	if g.window, err = sdl.CreateWindow("", sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED, g.windowWidth, g.windowHeight, sdl.WINDOW_SHOWN); err != nil {
		return fmt.Errorf("Error creating SDL_Window: %v", err)
	}

	// sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "linear")
	if g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return fmt.Errorf("Error creating SDL_Renderer: %v", err)
	}

	if err = g.renderer.SetLogicalSize(g.windowWidth, g.windowHeight); err != nil {
		return fmt.Errorf("Error setting renderer logical size: %v", err)
	}
	g.lastTime = sdl.GetTicks()
	return nil
}

func (g *Game) Close() {
	if g.assets != nil {
		g.assets.Close()
	}
	if g.renderer != nil {
		g.renderer.Destroy()
		g.renderer = nil
	}
	if g.window != nil {
		g.window.Destroy()
		g.window = nil
	}
	img.Quit()
	sdl.Quit()
	g.logger.Info("Game closed successfully!")
}
