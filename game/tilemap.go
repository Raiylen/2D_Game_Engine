package game

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type tile struct {
	col, row int
	id       int
}

type TilemapConfig struct {
	TileSize   int
	TileScale  int
	TileFormat int
}

type Tilemap struct {
	mapWidth  int
	mapHeight int
}

func LoadTilemap(path string, cfg TilemapConfig, fn func(tile)) (error, int, int) {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Error opening tilemap file: %v", err), 0, 0
	}
	defer f.Close()

	r := csv.NewReader(f)
	var row int
	var col int

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("Error reading tilemap file: %v", err), 0, 0
		}
		if col == 0 {
			col = len(record)
		}
		for col, s := range record {
			id, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				return fmt.Errorf("Error converting tilemap string to id at (%d, %d): %v", col, row, err), 0, 0
			}
			fn(tile{col: col, row: row, id: id})
		}
		row++
	}
	return nil, row, col
}
