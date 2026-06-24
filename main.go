package main

import (
	// "os"
	"github.com/raiylen/2d_game_engine/game"
	"github.com/raiylen/2d_game_engine/logger"
	// "os"
	// "runtime/pprof"
)

func main() {
	// f, _ := os.Create("cpu.prof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	log := logger.NewLogger()
	g := game.NewGame()

	defer g.Close()
	if err := g.Init(); err != nil {
		log.Err(err)
		return
	}

	if err := g.Setup(); err != nil {
		log.Err(err)
		return
	}

	g.Run()
}
