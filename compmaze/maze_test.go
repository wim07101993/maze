package compmaze

import (
	"testing"

	"github.com/wim07101993/maze"
	"github.com/wim07101993/maze/binmaze"
)

func MustContain(t *testing.T, m *Maze, cs []maze.Coordinate) {
	for _, c := range cs {
		if _, ok := m.Map[c]; !ok {
			t.Errorf("Coordinate %v not found", c)
		}
	}
}

func TestFromBinMaze(t *testing.T) {
	m, err := binmaze.FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}
	cm := FromBinMaze(m)

	if len(cm.Map) != 9 {
		t.Errorf("Length incorrect. Expected 9, got %v", len(cm.Map))
	}

	if cm.MaxX != 4 || cm.MaxY != 4 {
		t.Errorf("Dimensions are incorrect (%d, %d)", cm.MaxX, cm.MaxY)
	}

	cs := []maze.Coordinate{
		maze.Coordinate{X: 1, Y: 0},
		maze.Coordinate{X: 1, Y: 1},
		maze.Coordinate{X: 2, Y: 1},
		maze.Coordinate{X: 3, Y: 1},
		maze.Coordinate{X: 3, Y: 2},
		maze.Coordinate{X: 1, Y: 3},
		maze.Coordinate{X: 2, Y: 3},
		maze.Coordinate{X: 3, Y: 3},
		maze.Coordinate{X: 3, Y: 4},
	}

	MustContain(t, cm, cs)
}

func TestFromFile(t *testing.T) {
	m, err := FromFile("../mazes/5x5.maze")
	if err != nil {
		t.Error(err)
		return
	}

	if len(m.Map) != 9 {
		t.Errorf("Length incorrect. Expected 9, got %v", len(m.Map))
	}

	if m.MaxX != 4 || m.MaxY != 4 {
		t.Errorf("Dimensions are incorrect (%d, %d)", m.MaxX, m.MaxY)
	}

	cs := []maze.Coordinate{
		maze.Coordinate{X: 1, Y: 0},
		maze.Coordinate{X: 1, Y: 1},
		maze.Coordinate{X: 2, Y: 1},
		maze.Coordinate{X: 3, Y: 1},
		maze.Coordinate{X: 3, Y: 2},
		maze.Coordinate{X: 1, Y: 3},
		maze.Coordinate{X: 2, Y: 3},
		maze.Coordinate{X: 3, Y: 3},
		maze.Coordinate{X: 3, Y: 4},
	}

	MustContain(t, m, cs)
}
