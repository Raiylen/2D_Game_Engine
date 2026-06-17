package game

import (
	"fmt"

	log "github.com/raiylen/2d_game_engine/logger"

	// mgl "github.com/go-gl/mathgl/mgl32"
	img "github.com/veandco/go-sdl2/img"
	sdl "github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	isRunning    bool
	isFullscreen bool
	lastTime     uint32
	deltaTime    float32
	window       *sdl.Window
	windowWidth  int32
	windowHeight int32
	renderer     *sdl.Renderer
}

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

func (g *Game) Run() {
	frameDuration := uint32(1000 / framerate)

	for g.isRunning {
		frameStart := sdl.GetTicks()

		g.getInput()
		g.update()
		g.draw()

		elapsedTime := sdl.GetTicks() - frameStart
		if elapsedTime < frameDuration {
			sdl.Delay(frameDuration - elapsedTime)
		}
	}
}

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
					// g.renderer.SetLogicalSize(800, 600)

				}
			}
		}
	}
}

func (g *Game) update() {
	currentTime := sdl.GetTicks()
	g.deltaTime = float32(currentTime-g.lastTime) / 1000
	g.lastTime = currentTime
}

func (g *Game) draw() {
	g.renderer.SetDrawColor(21, 21, 21, 255)
	g.renderer.Clear()

	g.renderer.Present()
}

func (g *Game) Close() {
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
	log.Info("Game closed successfully!")
}

func (g *Game) LoadMedia() error {
	var err error

	return err
}

func NewGame() *Game {
	g := &Game{
		isRunning:    true,
		isFullscreen: false,
		windowWidth:  800,
		windowHeight: 600,
	}
	log.Info("New game created!")
	return g
}
