package memmaze

import (
	"bufio"
	"io"
	"os"

	"github.com/wim07101993/maze"
)

type Maze struct {
	width  int
	height int
	Exits  []*Tile
	Map    [][]*Tile
}

func FromReader(r io.Reader) *Maze {
	s := bufio.NewScanner(r)
	m := &Maze{Map: make([][]*Tile, 0)}

	for y := 0; s.Scan(); y++ {
		m.Map = append(m.Map, make([]*Tile, 0))
		bs := s.Bytes()

		for x, b := range bs {
			c := &Tile{}
			if maze.IsPath(b) {
				c.Type = Road
			} else {
				c.Type = Wall
			}

			c.X = x
			c.Y = y
			m.Map[y] = append(m.Map[y], c)
		}
	}

	m.width = len(m.Map[0])
	m.height = len(m.Map)

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

func (m *Maze) Range(f func(x, y int) bool) {
	for y := range m.Map {
		for x := range m.Map[y] {
			if f(x, y) {
				return
			}
		}
	}
}
