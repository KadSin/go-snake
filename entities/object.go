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
	Location    Coordinate
	ScreenStart Coordinate
	ScreenSize  Coordinate
	Shape       rune
	Direction   uint8
	Color       term.Attribute
}

func (object *Object) UpdateLocation(step int) error {
	switch object.Direction {
	case DIRECTION_UP:
		if object.Location.Y-step > object.ScreenStart.Y {
			object.Location.Y -= step
			return nil
		}
	case DIRECTION_RIGHT:
		if object.Location.X+step < object.ScreenSize.X {
			object.Location.X += step
			return nil
		}
	case DIRECTION_DOWN:
		if object.Location.Y+step < object.ScreenSize.Y {
			object.Location.Y += step
			return nil
		}
	case DIRECTION_LEFT:
		if object.Location.X-step > object.ScreenStart.X {
			object.Location.X -= step
			return nil
		}
	}

	return errors.New("Object exceeds the boundary")
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
