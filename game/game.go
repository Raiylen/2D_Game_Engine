package game

import (
	"fmt"

	assets "github.com/raiylen/2d_game_engine/assets"
	components "github.com/raiylen/2d_game_engine/engine/components"
	ecs "github.com/raiylen/2d_game_engine/engine/ecs"
	systems "github.com/raiylen/2d_game_engine/engine/systems"
	logger "github.com/raiylen/2d_game_engine/logger"

	// mgl "github.com/go-gl/mathgl/mgl32"
	img "github.com/veandco/go-sdl2/img"
	sdl "github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	isRunning    bool
	isFullscreen bool
	lastTime     uint32
	deltaTime    float64
	window       *sdl.Window
	windowWidth  int32
	windowHeight int32
	renderer     *sdl.Renderer
	world        *ecs.World
	assets       *assets.AssetManager
	logger       *logger.Logger
	tilemap      *Tilemap
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
				case sdl.SCANCODE_F3:
					g.world.Flags["debug"] = !g.world.Flags["debug"]
				}
			}
		}
	}
}

func (g *Game) update() {
	currentTime := sdl.GetTicks()
	g.deltaTime = float64(currentTime-g.lastTime) / 1000
	g.lastTime = currentTime

	g.world.Update(g.deltaTime)
}

func (g *Game) draw() {
	g.renderer.SetDrawColor(21, 21, 21, 255)
	g.renderer.Clear()

	g.world.Render()

	g.renderer.Present()
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

func (g *Game) Setup() error {
	g.world = ecs.NewWorld()
	g.world.SetLogger(g.logger)
	g.assets = assets.NewAssetManager(g.renderer)

	g.world.RegisterSystem(systems.NewDamageSystem())
	g.world.RegisterSystem(systems.NewMovementSystem())
	g.world.RegisterSystem(systems.NewAnimationSystem())
	g.world.RegisterSystem(systems.NewCollisionSystem())
	g.world.RegisterSystem(systems.NewRenderSystem(g.renderer, g.assets))
	g.world.RegisterSystem(systems.NewDebugRenderSystem(g.renderer))

	return g.loadLevel()
}

func (g *Game) loadLevel() error {
	var err error

	if err = g.assets.AddTexture("tank-image", "assets/images/tank-panther-right.png"); err != nil {
		g.logger.Warn(err.Error())
	}
	if err = g.assets.AddTexture("truck-image", "assets/images/truck-ford-down.png"); err != nil {
		g.logger.Warn(err.Error())
	}
	if err = g.assets.AddTexture("tilemap-image", "assets/tilemaps/jungle.png"); err != nil {
		g.logger.Warn(err.Error())
	}
	if err = g.assets.AddTexture("chopper-image", "assets/images/chopper.png"); err != nil {
		g.logger.Warn(err.Error())
	}

	cfg := TilemapConfig{
		TileSize:   32,
		TileScale:  2,
		TileFormat: 10,
	}

	err = LoadTilemap("assets/tilemaps/jungle.map", cfg, func(t tile) {
		e := g.world.NewEntity()
		ecs.AddComponent(g.world.Registry, e, components.Transform{
			X:     float64((t.col * cfg.TileSize) * cfg.TileScale),
			Y:     float64((t.row * cfg.TileSize) * cfg.TileScale),
			Scale: float64(cfg.TileScale),
		})
		ecs.AddComponent(g.world.Registry, e, components.Sprite{
			AssetID:    "tilemap-image",
			Width:      cfg.TileSize,
			Height:     cfg.TileSize,
			SourcePosX: (t.id % cfg.TileFormat) * cfg.TileSize,
			SourcePosY: (t.id / cfg.TileFormat) * cfg.TileSize,
		})
	})
	if err != nil {
		g.logger.Warn(err.Error())
	}

	tank := g.world.NewEntity()
	ecs.AddComponent(g.world.Registry, tank, components.Transform{X: 500, Y: 20, Scale: 1})
	ecs.AddComponent(g.world.Registry, tank, components.RigidBody{DX: -50, DY: 0})
	ecs.AddComponent(g.world.Registry, tank, components.Sprite{AssetID: "tank-image", Width: 32, Height: 32, Layer: 12})
	ecs.AddComponent(g.world.Registry, tank, components.BoxCollider{Width: 32, Height: 32})

	truck := g.world.NewEntity()
	ecs.AddComponent(g.world.Registry, truck, components.Transform{X: 20, Y: 20, Scale: 1.0})
	ecs.AddComponent(g.world.Registry, truck, components.RigidBody{DX: 50, DY: 0})
	ecs.AddComponent(g.world.Registry, truck, components.Sprite{AssetID: "truck-image", Width: 32, Height: 32, Layer: 11})
	ecs.AddComponent(g.world.Registry, truck, components.BoxCollider{Width: 32, Height: 32})

	chopper := g.world.NewEntity()
	ecs.AddComponent(g.world.Registry, chopper, components.Transform{X: 200, Y: 200, Scale: 1.0})
	ecs.AddComponent(g.world.Registry, chopper, components.RigidBody{DX: 10, DY: 0})
	ecs.AddComponent(g.world.Registry, chopper, components.Sprite{AssetID: "chopper-image", Width: 32, Height: 32, Layer: 11})
	ecs.AddComponent(g.world.Registry, chopper, components.Animation{NumFrames: 2, CurrentFrame: 1, FrameRate: 10, Loop: true})

	return nil
}

func NewGame() *Game {
	g := &Game{
		isRunning:    true,
		isFullscreen: false,
		windowWidth:  1280,
		windowHeight: 720,
		logger:       logger.NewLogger(),
	}
	g.logger.Info("New game created!")
	return g
}
