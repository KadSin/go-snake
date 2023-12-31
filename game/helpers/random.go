package helpers

import "math/rand"

func RandomIntElement(first int, second int) int {
	if RandomBoolean() {
		return first
	} else {
		return second
	}
}

func RandomNumberBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}

func RandomBoolean() bool {
	return rand.Float32() > 0.5
}
