package generator

import (
	"fmt"

	"github.com/wim07101993/maze"
)

type Maze struct {
	width    int
	height   int
	x        int
	y        int
	value    byte
	solution maze.Path
	Map      [][]int
}

func New(width, height int) *Maze {
	return &Maze{
		width:    width,
		height:   height,
		solution: maze.Path{},
		Map:      make([][]int, 0),
	}
}

func (m *Maze) GeneratePath() maze.Path {
	m.solution = GeneratePath(m.width, m.height)
	return m.solution
}

func (m *Maze) MergeMapAndSolution() [][]int {
	if m.Map == nil {
		m.GenerateMap()
	}
	if m.solution == nil {
		m.GeneratePath()
	}

	for c := range m.solution {
		fmt.Println(c)
	}

	return m.Map
}

func (m *Maze) GenerateMap() [][]int {
	is := make([][]int, m.height)

	for y := range is {
		is[y] = m.GenerateRow()
	}

	m.Map = is
	return is
}

func (m *Maze) GenerateRow() []int {
	blocks := m.width/64 + 1
	is := make([]int, blocks)

	for i := 0; i < blocks; i++ {
		g1 := rnd.GenInt()
		g2 := rnd.GenInt()
		is[i] = g1 | g2
	}

	for i := 0; i < 64-m.width%64; i++ {
		is[0] = is[0] &^ (1 << uint(i))
	}

	return is
}
