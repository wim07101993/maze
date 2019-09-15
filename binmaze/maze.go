package binmaze

import (
	"bufio"
	"io"
	"os"

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
