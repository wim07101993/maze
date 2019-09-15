package binmaze

import (
	"bufio"
	"io"
	"os"
	"strings"

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
