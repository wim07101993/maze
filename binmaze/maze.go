package binmaze

import (
	"bufio"
	"io"
)

type Maze [][]bool

func FromReader(r io.Reader) Maze {
	s := bufio.NewScanner(r)
}
