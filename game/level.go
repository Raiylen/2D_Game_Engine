package game

import (
	components "github.com/raiylen/2d_game_engine/engine/components"
	ecs "github.com/raiylen/2d_game_engine/engine/ecs"
)

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
	if err = g.assets.AddTexture("chopper-image", "assets/images/chopper-spritesheet.png"); err != nil {
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
	ecs.AddComponent(g.world.Registry, chopper, components.RigidBody{DX: 0, DY: 0})
	ecs.AddComponent(g.world.Registry, chopper, components.Sprite{AssetID: "chopper-image", Width: 32, Height: 32, Layer: 11})
	ecs.AddComponent(g.world.Registry, chopper, components.Animation{NumFrames: 2, CurrentFrame: 1, FrameRate: 10, Loop: true})
	ecs.AddComponent(g.world.Registry, chopper, components.KeyboardControl{UpVelocity: 30, DownVelocity: 30, RightVelocity: 30, LeftVelocity: 30})

	return nil
}
