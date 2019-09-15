package memmaze

type Maze struct {
	width  int
	height int
	Exits  []*Tile
	Map    [][]*Tile
}
