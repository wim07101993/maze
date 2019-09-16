package memmaze

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"

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

func (m *Maze) Explore(at *Tile, p maze.Path, out chan<- []maze.Path) {
	explorers := make([]chan []maze.Path, len(maze.PossibleDirections))

	for _, d := range maze.PossibleDirections {
		t, c := m.Look(d, at, p)

		explorer := make(chan []maze.Path)
		explorers[d] = explorer

		newP := make(maze.Path, len(p))
		copy(newP, p)
		newP = append(newP, &at.Coordinate)

		switch t {
		case Road:
			go m.Explore(c, newP, explorer)
		case Wall, Duplicate:
			go func() {
				explorer <- nil
				close(explorer)
			}()
		case Exit:
			go func() {
				explorer <- []maze.Path{newP}
				close(explorer)
			}()
		}

	}

	out <- m.WaitForExplorers(explorers)
	close(out)
}

func (m *Maze) WaitForExplorers(explorers []chan []maze.Path) []maze.Path {
	ret := make([]maze.Path, 0)
	merge := make(chan []maze.Path)

	var wg sync.WaitGroup
	wg.Add(len(explorers))

	for _, e := range explorers {
		go func(c <-chan []maze.Path, merge chan<- []maze.Path) {
			merge <- <-c
			wg.Done()
		}(e, merge)
	}

	go func() {
		wg.Wait()
		close(merge)
	}()

	for p := range merge {
		if p != nil {
			ret = append(ret, p...)
		}
	}

	return ret
}

func (m *Maze) Look(to maze.Direction, at *Tile, p maze.Path) (TileType, *Tile) {
	var x, y int

	switch to {
	case maze.North:
		x = at.X
		y = at.Y - 1
	case maze.East:
		x = at.X + 1
		y = at.Y
	case maze.South:
		x = at.X
		y = at.Y + 1
	case maze.West:
		x = at.X - 1
		y = at.Y
	}

	if y >= m.height || x >= m.width || y < 0 || x < 0 {
		return Exit, nil
	}

	newT := m.Map[y][x]

	if newT.Type == Wall {
		return Wall, newT
	}

	if p.Contains(&newT.Coordinate) {
		return Duplicate, newT
	}

	return Road, newT
}

func (m *Maze) String() string {
	builder := strings.Builder{}

	for y := range m.Map {
		for x := range m.Map[y] {
			switch m.Map[y][x].Type {
			case Road, Exit:
				builder.WriteString(".")
			case Wall:
				builder.WriteString("x")
			case Duplicate:
				builder.WriteString("d")
			}
		}
		builder.WriteString("\r\n")
	}

	return builder.String()
}
