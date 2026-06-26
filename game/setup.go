package game

import (
	assets "github.com/raiylen/2d_game_engine/assets"
	ecs "github.com/raiylen/2d_game_engine/engine/ecs"
	systems "github.com/raiylen/2d_game_engine/engine/systems"
	sdl "github.com/veandco/go-sdl2/sdl"
)

func (g *Game) Setup() error {
	g.world = ecs.NewWorld()
	g.assets = assets.NewAssetManager(g.renderer)
	g.world.SetLogger(g.logger)
	g.keystate = sdl.GetKeyboardState()

	g.world.RegisterSystem(systems.NewDamageSystem())
	g.world.RegisterSystem(systems.NewMovementSystem())
	g.world.RegisterSystem(systems.NewAnimationSystem())
	g.world.RegisterSystem(systems.NewCollisionSystem())
	g.world.RegisterSystem(systems.NewRenderSystem(g.renderer, g.assets))
	g.world.RegisterSystem(systems.NewDebugRenderSystem(g.renderer))
	g.world.RegisterSystem(systems.NewControlSystem())

	return g.loadLevel()
}
