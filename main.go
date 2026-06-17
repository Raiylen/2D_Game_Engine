package main

import (
	// "fmt"
	// "os"

	game "github.com/raiylen/2d_game_engine/game"
	log "github.com/raiylen/2d_game_engine/logger"
)

func main() {
	g := game.NewGame()

	defer g.Close()
	if err := g.Init(); err != nil {
		log.Err(err)
		return
	}

	if err := g.LoadMedia(); err != nil {
		log.Err(err)
		return
	}

	g.Run()
}
