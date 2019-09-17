package compmaze

import (
	"bufio"
	"io"
	"os"

	"github.com/wim07101993/maze"
	"github.com/wim07101993/maze/binmaze"
)

type Maze struct {
	Map  map[maze.Coordinate]*Tile
	MaxX int
	MaxY int
}

func FromBinMaze(m binmaze.Maze) *Maze {
	c := &Maze{
		Map:  make(map[maze.Coordinate]*Tile),
		MaxX: len(m[0]) - 1,
		MaxY: len(m) - 1,
	}

	m.Range(func(x, y int) bool {
		if !m[y][x] {
			return false
		}

		t := &Tile{}
		t.X = x
		t.Y = y
		c.add(t)
		return false
	})

	return c
}

func FromReader(r io.Reader) *Maze {
	s := bufio.NewScanner(r)
	m := &Maze{Map: make(map[maze.Coordinate]*Tile)}

	maxX := 0
	x := -1
	y := -1
	for s.Scan() {
		y++
		bs := s.Bytes()
		for _, b := range bs {
			x++
			if !maze.IsPath(b) {
				continue
			}

			t := &Tile{}
			t.X = x
			t.Y = y

			m.add(t)
		}
		if maxX < x {
			maxX = x
		}
		x = -1
	}

	m.MaxX = maxX
	m.MaxY = y

	return m
}

func FromFile(path string) (*Maze, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return FromReader(file), nil
}

func (m *Maze) add(t *Tile) {
	// north of the current tile is another road-tile
	if north, ok := m.Map[maze.Coordinate{X: t.X, Y: t.Y - 1}]; ok {
		t.North = north
		north.South = t
	}
	if east, ok := m.Map[maze.Coordinate{X: t.X + 1, Y: t.Y}]; ok {
		t.East = east
		east.West = t
	}
	if south, ok := m.Map[maze.Coordinate{X: t.X, Y: t.Y + 1}]; ok {
		t.South = south
		south.North = t
	}
	if west, ok := m.Map[maze.Coordinate{X: t.X - 1, Y: t.Y}]; ok {
		t.West = west
		west.East = t
	}

	m.Map[t.Coordinate] = t
}

func (m *Maze) Range(f func(k maze.Coordinate, v *Tile) bool) {
	for k, v := range m.Map {
		if f(k, v) {
			return
		}
	}
}
