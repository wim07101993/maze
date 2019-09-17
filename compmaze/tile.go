package compmaze

import "github.com/wim07101993/maze"

type Tile struct {
	maze.Coordinate
	North *Tile
	East  *Tile
	South *Tile
	West  *Tile
}
