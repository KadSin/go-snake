package entities

type Coordinate struct {
	X, Y int
}

type Screen struct {
	Start Coordinate
	End   Coordinate
}
