package game

import (
	assets "github.com/raiylen/2d_game_engine/assets"
	ecs "github.com/raiylen/2d_game_engine/engine/ecs"
	logger "github.com/raiylen/2d_game_engine/logger"

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
	keystate     []uint8
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
