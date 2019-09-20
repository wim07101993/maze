package generator

import (
	"testing"

	"github.com/wim07101993/maze"
)

func TestCheckCanGoDirection(t *testing.T) {
	d, ok := checkCanGoDirection(maze.South, maze.South, maze.South)
	if !ok {
		t.Error("Can go south after going south twice")
	} else if d != maze.South {
		t.Errorf("Expected to go south, not %v", d)
	}

	d, ok = checkCanGoDirection(maze.East, maze.South, maze.South)
	if !ok {
		t.Error("Can go east after going south twice")
	} else if d != maze.East {
		t.Errorf("Expected to go east, not %v", d)
	}

	d, ok = checkCanGoDirection(maze.North, maze.South, maze.South)
	if !ok {
		t.Error("Can go vertical after going south twice")
	} else if d != maze.South {
		t.Errorf("Expected to go sourth, not %v", d)
	}

	_, ok = checkCanGoDirection(maze.North, maze.East, maze.South)
	if ok {
		t.Error("Can not go north after going south and then east")
	}

	_, ok = checkCanGoDirection(maze.North, maze.West, maze.South)
	if ok {
		t.Error("Cannot go north after going south and then west")
	}
}
