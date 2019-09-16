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
	m.FindExits()

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

func (m *Maze) FindExits() []*Tile {
	ts := make([]*Tile, 0)

	for _, t := range m.Map[0] {
		if t.Type == Road {
			t.Type = Exit
			ts = append(ts, t)
		}
	}

	for _, t := range m.Map[m.height-1] {
		if t.Type == Road {
			t.Type = Exit
			ts = append(ts, t)
		}
	}

	for _, r := range m.Map {
		if r[0].Type == Road {
			r[0].Type = Exit
			ts = append(ts, r[0])
		}
		if r[m.width-1].Type == Road {
			r[m.width-1].Type = Exit
			ts = append(ts, r[m.width-1])
		}
	}

	return ts
}
