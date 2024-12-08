package state

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func First() Direction {
	return Up
}

func Next(d Direction) Direction {
	switch d {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return Up
	}
}
