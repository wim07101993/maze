package binmaze

import (
	"bufio"
	"io"

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
