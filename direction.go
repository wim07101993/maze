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
