package binmaze

import (
	"errors"
	"fmt"
	"testing"

	"github.com/wim07101993/maze"
)

func CompareMaze(m Maze, bs [][]bool) error {
	var err error
	m.Range(func(x, y int) bool {
		if m[y][x] != bs[y][x] {
			if bs[y][x] {
				err = errors.New(fmt.Sprintf("(%d, %d) is not a wall", x, y))
			} else {
				err = errors.New(fmt.Sprintf("(%d, %d) is a wall", x, y))
			}
			return true
		}
		return false
	})

	return err
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

func TestGetOpenDirections(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	ds := m.GetOpenDirections(1, 0)
	if !ds.Contains(maze.North) || !ds.Contains(maze.South) {
		t.Errorf("Did not get the correct direction: %v", ds)
	}
	ds = m.GetOpenDirections(1, 1)
	if !ds.Contains(maze.North) || !ds.Contains(maze.East) {
		t.Errorf("Did not get the correct direction: %v", ds)
	}
	ds = m.GetOpenDirections(3, 3)
	if !ds.Contains(maze.North) || !ds.Contains(maze.West) || !ds.Contains(maze.South) {
		t.Errorf("Did not get the correct directions: %v", ds)
	}
	ds = m.GetOpenDirections(1, -1)
	if !ds.Contains(maze.North) || !ds.Contains(maze.South) {
		t.Errorf("Did not get the correct directions: %v", ds)
	}
}

func TestNormalize(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	x, y := m.normalize(-1, -1)
	if x != 0 || y != 0 {
		t.Errorf("Exprected, 0,0, got %d,%d", x, y)
	}
	x, y = m.normalize(1, 1)
	if x != 1 || y != 1 {
		t.Errorf("Exprected, 1,1, got %d,%d", x, y)
	}
	x, y = m.normalize(3, 20)
	if x != 3 || y != 4 {
		t.Errorf("Exprected, 3,4, got %d,%d", x, y)
	}
	x, y = m.normalize(5, 4)
	if x != 4 || y != 4 {
		t.Errorf("Exprected, 4,4, got %d,%d", x, y)
	}
}

func TestMove(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
	}

	x, y := m.move(-1, -1, maze.North)
	if x != 0 || y != 0 {
		t.Errorf("Expected 0,0, got %d,%d", x, y)
	}
	x, y = m.move(1, 1, maze.East)
	if x != 2 || y != 1 {
		t.Errorf("Expected 2,1, got %d,%d", x, y)
	}
	x, y = m.move(2, 1, maze.South)
	if x != 2 || y != 2 {
		t.Errorf("Expected 2,2, got %d,%d", x, y)
	}
	x, y = m.move(3, 3, maze.North)
	if x != 3 || y != 2 {
		t.Errorf("Expected 3,2, got %d,%d", x, y)
	}
	x, y = m.move(3, 3, maze.West)
	if x != 2 || y != 3 {
		t.Errorf("Expected 2,3, got %d,%d", x, y)
	}
}

func TestRemoveDeadEnds4x4(t *testing.T) {
	m, err := FromFile("../mazes/4x4.maze")
	if err != nil {
		t.Error(err)
		return
	}

	m.RemoveDeadEnds()

	bs := [][]bool{
		{false, true, false, false},
		{false, true, true, false},
		{false, false, true, false},
		{false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestRemoveDeadEnds5x5(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	m.RemoveDeadEnds()

	bs := [][]bool{
		{false, true, false, false, false},
		{false, true, true, true, false},
		{false, false, false, true, false},
		{false, false, false, true, false},
		{false, false, false, true, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestRemoveDeadEnds6x6(t *testing.T) {
	m, err := FromFile("../mazes/6x6.maze")
	if err != nil {
		t.Error(err)
		return
	}

	m.RemoveDeadEnds()

	bs := [][]bool{
		{false, true, false, false, false, false},
		{false, true, false, false, false, false},
		{false, true, true, false, false, false},
		{false, false, true, false, false, false},
		{false, false, true, true, false, false},
		{false, false, false, true, false, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}

func TestRemoveDeadEnds9x7(t *testing.T) {
	m, err := FromFile("../mazes/9x7.maze")
	if err != nil {
		t.Error(err)
		return
	}

	m.RemoveDeadEnds()

	bs := [][]bool{
		{false, true, false, false, false, false, false, false, false},
		{false, true, false, false, true, true, true, true, false},
		{false, true, true, false, true, false, false, true, false},
		{false, false, true, true, true, false, false, true, false},
		{false, false, true, false, false, false, false, true, false},
		{false, false, true, true, true, true, true, true, false},
		{false, false, true, false, false, true, false, false, false},
	}

	if err := CompareMaze(m, bs); err != nil {
		fmt.Println(m)
		t.Error(err)
	}
}
