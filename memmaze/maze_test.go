package memmaze

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/wim07101993/maze"
)

func CompareMaze(m *Maze, bs [][]bool) error {
	var err error
	m.Range(func(x, y int) bool {
		switch m.Map[y][x].Type {
		case Road, Exit:
			if !bs[y][x] {
				err = errors.New(fmt.Sprintf("(%d, %d) => (%v) is not a wall", x, y, m.Map[y][x].Type))
			}
		case Wall:
			if bs[y][x] {
				err = errors.New(fmt.Sprintf("(%d, %d) => (%v) is a wall", x, y, m.Map[y][x].Type))
			}
		case Duplicate:
			err = errors.New(fmt.Sprintf("Duplicate (%d, %d)\r\n", x, y))
		}
		return false
	})

	return err
}

func PrintBools(bs [][]bool) {
	builder := strings.Builder{}

	for y := range bs {
		for x := range bs[y] {
			if bs[y][x] {
				builder.WriteString(".")
			} else {
				builder.WriteString("x")
			}
		}
		builder.WriteString("\r\n")
	}

	fmt.Println(builder.String())
}

func TestParse4x4(t *testing.T) {
	m, err := FromFile("../mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false},
		{false, true, true, false},
		{false, false, true, false},
		{false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		PrintBools(bs)
		fmt.Println(m)
		t.Error(err)
	}
}

func TestParse5x5(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestParse6x6(t *testing.T) {
	m, err := FromFile("../mazes/6x6.maze")
	if err != nil {
		t.Error(err)
		return
	}

	bs := [][]bool{
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, true, true, false, false},
		{false, false, true, false, true, false},
		{false, false, true, true, true, false},
		{false, false, false, true, false, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestExplore4x4(t *testing.T) {
	m, err := FromFile("../mazes/4x4.maze")
	if err != nil {
		t.Error(err)
	}
	out := make(chan []maze.Path)

	go m.Explore(m.Map[0][1], maze.Path{}, out)

	paths := <-out
	p := paths[1]
	if len(p) != 5 {
		fmt.Println(len(p))
		fmt.Println(p)
		t.Errorf("The length of the route should be 5, not %d", len(p))
	}
}
