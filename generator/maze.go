package generator

import "github.com/wim07101993/maze"

type Maze struct {
	width    int
	height   int
	x        int
	y        int
	value    byte
	solution maze.Path
}

func New(width, height int) *Maze {
	return &Maze{
		width:  width,
		height: height,
	}
}

func (m *Maze) generateSolution() {
	m.solution = maze.GeneratePath(m.width-1, m.height-1)
}
