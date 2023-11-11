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

type Object struct {
	X         int
	Y         int
	Shape     rune
	Direction uint8
	Color     term.Attribute
}

var TerminalSize = term.Size

func (object *Object) UpdateLocation(step int) error {
	var width, height = TerminalSize()

	switch object.Direction {
	case DIRECTION_UP:
		if object.Y > 0 {
			object.Y -= step
			return nil
		}
	case DIRECTION_RIGHT:
		if object.X < width-1 {
			object.X += step
			return nil
		}
	case DIRECTION_DOWN:
		if object.Y < height-1 {
			object.Y += step
			return nil
		}
	case DIRECTION_LEFT:
		if object.X > 0 {
			object.X -= step
			return nil
		}
	}

	return errors.New("Object is on the boundry")
}

func (object *Object) MoveUp() {
	object.Direction = DIRECTION_UP
}

func (object *Object) MoveRight() {
	object.Direction = DIRECTION_RIGHT
}

func (object *Object) MoveDown() {
	object.Direction = DIRECTION_DOWN
}

func (object *Object) MoveLeft() {
	object.Direction = DIRECTION_LEFT
}
