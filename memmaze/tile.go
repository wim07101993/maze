package memmaze

import (
	"github.com/wim07101993/maze"
)

type TileType int

const (
	Wall TileType = iota
	Road
	Exit
	Duplicate
)

type Tile struct {
	maze.Coordinate
	Type TileType
}

func (t TileType) String() string {
	switch t {
	case Wall:
		return "Wall"
	case Road:
		return "Road"
	case Exit:
		return "Exit"
	case Duplicate:
		return "Duplicate"
	default:
		return "Unknown"
	}
}
