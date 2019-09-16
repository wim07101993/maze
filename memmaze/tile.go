package memmaze

import (
	"github.com/wim07101993/maze"
)

type TileType int

const (
	Wall TileType = iota
	Road
	Exit
)

type Tile struct {
	maze.Coordinate
	Type TileType
}
