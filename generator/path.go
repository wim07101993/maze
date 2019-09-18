package generator

import (
	"github.com/wim07101993/maze"
)

var (
	rnd *Randomizer
)

func init() {
	rnd = NewRandomizer()
}

func GeneratePath(maxX, maxY int) maze.Path {
	var startX, x, y int
	for x == 0 || x == maxX {
		x = rnd.GenInt()
		startX = x
	}
	y = 1

	d1 := maze.South
	d2 := maze.South
	p := maze.Path{
		&maze.Coordinate{X: startX, Y: 0},
		&maze.Coordinate{X: startX, Y: 1},
	}

	for {
		d := chooseDirection(d1, d2)

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

		p = append(p, &maze.Coordinate{X: x, Y: y})

		// if we are back at the start, keep on going
		if y == 0 && x == startX {
			continue
		}

		// if we reached a border, finish
		if x == 0 || x == maxX ||
			y == 0 || y == maxY {
			return p
		}
	}
}

func chooseDirection(d1, d2 maze.Direction) maze.Direction {
	d := maze.Direction(rnd.Gen(2))

	d1IsVertical := (d1 != maze.North && d1 != maze.South)
	for {

		// if the direction is oposit, go in the same way
		if ((d == maze.North || d == maze.South) && d1IsVertical) ||
			((d == maze.East || d == maze.West) && !d1IsVertical) {
			return d1
		}

		// make sure we do not walk next to already taken paths
		// xx xx      xx xx
		// xx xx  -\  xx  x
		// xx  x  -/  xx  x
		// xxxxx      xxxxx
		if (d == maze.North && (d2 != maze.South || !d1IsVertical)) ||
			(d == maze.East && (d2 != maze.West || d1IsVertical)) ||
			(d == maze.South && (d2 != maze.North || !d1IsVertical)) ||
			(d == maze.West && (d2 != maze.East || d1IsVertical)) {
			return d
		}

		d = maze.Direction(rnd.Gen(2))
	}
}
