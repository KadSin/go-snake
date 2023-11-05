package entities

import (
	"errors"

	term "github.com/nsf/termbox-go"
)

const (
	DIRECTION_UP = iota
	DIRECTION_RIGHT
	DIRECTION_DOWN
	DIRECTION_LEFT
)

type Animal struct {
	X         int
	Y         int
	Shape     rune
	Direction int
}

func (animal *Animal) UpdateLocation() error {
	var width, height = term.Size()

	switch animal.Direction {
	case DIRECTION_UP:
		if animal.Y > 0 {
			animal.Y--
			return nil
		}
	case DIRECTION_RIGHT:
		if animal.X < width-1 {
			animal.X++
			return nil
		}
	case DIRECTION_DOWN:
		if animal.Y < height-1 {
			animal.Y++
			return nil
		}
	case DIRECTION_LEFT:
		if animal.X > 0 {
			animal.X--
			return nil
		}
	}

	return errors.New("Animal is on the boundry")
}

func (animal *Animal) MoveUp() {
	animal.Direction = DIRECTION_UP
}

func (animal *Animal) MoveRight() {
	animal.Direction = DIRECTION_RIGHT
}

func (animal *Animal) MoveDown() {
	animal.Direction = DIRECTION_DOWN
}

func (animal *Animal) MoveLeft() {
	animal.Direction = DIRECTION_LEFT
}
