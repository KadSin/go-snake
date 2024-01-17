package helpers

import (
	"kadsin/shoot-run/game/assets"
	"math/rand"
)

func RandomIntElement(first int, second int) int {
	if RandomBoolean() {
		return first
	} else {
		return second
	}
}

func RandomNumberBetween(min int, max int) int {
	return rand.Intn(max+1-min) + min
}

func RandomBoolean() bool {
	return rand.Float32() > 0.5
}

func RandomCoordinateOnBorders(screen assets.Screen) assets.Coordinate {
	coordinate := assets.Coordinate{
		X: RandomIntElement(screen.Start.X, screen.End.X),
		Y: RandomIntElement(screen.Start.Y, screen.End.Y),
	}

	if RandomBoolean() {
		coordinate.X = RandomNumberBetween(screen.Start.X, screen.End.X)
	} else {
		coordinate.Y = RandomNumberBetween(screen.Start.Y, screen.End.Y)
	}

	return coordinate
}

func RandomCoordinate(screen assets.Screen, distance assets.Coordinate) assets.Coordinate {
	return assets.Coordinate{
		X: RandomNumberBetween(screen.Start.X+distance.X, screen.End.X-distance.X),
		Y: RandomNumberBetween(screen.Start.Y+distance.Y, screen.End.Y-distance.Y),
	}
}
