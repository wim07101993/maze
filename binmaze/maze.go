package binmaze

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/wim07101993/maze"
)

type Maze [][]bool

func FromReader(r io.Reader) Maze {
	s := bufio.NewScanner(r)

	m := make(Maze, 0)
	for y := 0; s.Scan(); y++ {
		m = append(m, make([]bool, 0))
		bs := s.Bytes()
		for _, b := range bs {
			m[y] = append(m[y], maze.IsPath(b))
		}
	}

	return m
}

func FromFile(path string) (Maze, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return FromReader(file), nil
}

func (m Maze) Range(f func(x, y int) bool) {
	for y := range m {
		for x := range m[y] {
			if f(x, y) {
				return
			}
		}
	}
}

func (m Maze) GetOpenDirections(x, y int) maze.Directions {
	x, y = m.normalize(x, y)
	ds := make(maze.Directions, 0)

	if x == 0 || m[y][x-1] {
		ds = append(ds, maze.West)
	}
	if x == len(m[0])-1 || m[y][x+1] {
		ds = append(ds, maze.East)
	}
	if y == 0 || m[y-1][x] {
		ds = append(ds, maze.North)
	}
	if y == len(m)-1 || m[y+1][x] {
		ds = append(ds, maze.South)
	}

	return ds
}

func (m Maze) RemoveDeadEnds() {
	// slice of channels to indicate whether all go routines are finished
	cs := make([]chan bool, 0)

	// resolve dead ends for all tiles
	m.Range(func(x, y int) bool {
		c := make(chan bool)
		cs = append(cs, c)
		m.resolveDeadEnds(x, y)
		return false
	})
}

func (m Maze) resolveDeadEnds(x, y int) {
	maxX := x
	maxY := y

	for {
		// current is wall
		if !m[y][x] {
			return
		}

		ds := m.GetOpenDirections(x, y)

		switch len(ds) {
		// no exits
		case 0:
			m[y][x] = false
			return

		// dead end
		case 1:
			m[y][x] = false
			x, y = m.move(x, y, ds[0])
			// if further than starting point, return
			if y > maxY || (y == maxY && x > maxX) {
				return
			}

		// has entrance and exit
		default:
			return
		}
	}
}

func (m Maze) move(x, y int, d maze.Direction) (newX, newY int) {
	switch d {
	case maze.North:
		y--
	case maze.East:
		x++
	case maze.South:
		y++
	case maze.West:
		x--
	}

	return m.normalize(x, y)
}

func (m Maze) normalize(x, y int) (newX, newY int) {
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	if x > len(m[0])-1 {
		x = len(m[0]) - 1
	}
	if y > len(m)-1 {
		y = len(m) - 1
	}
	return x, y
}

func (m Maze) String() string {
	builder := strings.Builder{}

	for y := range m {
		for x := range m[y] {
			if m[y][x] {
				builder.WriteString(".")
			} else {
				builder.WriteString("x")
			}
		}
		builder.WriteString("\r\n")
	}

	return builder.String()
}

func Wait(cs []chan bool) {
	merge := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan bool, merge chan<- bool) {
			for b := range c {
				merge <- b
			}
			wg.Done()
		}(c, merge)
	}

	wg.Wait()
	close(merge)
}
