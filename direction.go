package maze

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Directions []Direction

var (
	PossibleDirections Directions
)

func init() {
	PossibleDirections = make(Directions, 4, 4)
	PossibleDirections[North] = North
	PossibleDirections[East] = East
	PossibleDirections[South] = South
	PossibleDirections[West] = West
}

func (ds Directions) Contains(find Direction) bool {
	for _, d := range ds {
		if d == find {
			return true
		}
	}
	return false
}

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return ""
	}
}
