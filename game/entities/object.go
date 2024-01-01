package entities

import (
	"errors"
	"kadsin/shoot-run/game/assets"

	term "github.com/nsf/termbox-go"
)

const (
	DIRECTION_UP = iota + 1
	DIRECTION_RIGHT
	DIRECTION_DOWN
	DIRECTION_LEFT
)

type Object struct {
	Location            assets.Coordinate
	Screen              assets.Screen
	Shape               rune
	Direction           uint8
	AdditionalDirection uint8
	Color               term.Attribute
}

func (object *Object) UpdateLocation(step int) error {
	nextStep := object.NextStep(step)

	if nextStep == object.Location {
		return errors.New("Object exceeds the boundary")
	} else {
		object.Location = nextStep

		return nil
	}
}

func (object Object) NextStep(step int) assets.Coordinate {
	object.Location = object.refineNextStepByDirection(object.Direction, step)
	object.Location = object.refineNextStepByDirection(object.AdditionalDirection, step)

	return object.Location
}

func (object Object) refineNextStepByDirection(direction uint8, step int) assets.Coordinate {
	switch direction {
	case DIRECTION_UP:
		if object.Location.Y-step >= object.Screen.Start.Y {
			object.Location.Y -= step
		}
	case DIRECTION_RIGHT:
		if object.Location.X+step < object.Screen.End.X {
			object.Location.X += step
		}
	case DIRECTION_DOWN:
		if object.Location.Y+step < object.Screen.End.Y {
			object.Location.Y += step
		}
	case DIRECTION_LEFT:
		if object.Location.X-step >= object.Screen.Start.X {
			object.Location.X -= step
		}
	}

	return object.Location
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

func (object *Object) AdditionalMoveUp() {
	object.AdditionalDirection = DIRECTION_UP
}

func (object *Object) AdditionalMoveRight() {
	object.AdditionalDirection = DIRECTION_RIGHT
}

func (object *Object) AdditionalMoveDown() {
	object.AdditionalDirection = DIRECTION_DOWN
}

func (object *Object) AdditionalMoveLeft() {
	object.AdditionalDirection = DIRECTION_LEFT
}

func (object *Object) DoesHit(target Object) bool {
	isObjectNearbyTargetHorizontal := object.Location.X >= target.Location.X-1 && object.Location.X <= target.Location.X+1
	isObjectNearbyTargetVertical := object.Location.Y >= target.Location.Y-1 && object.Location.Y <= target.Location.Y+1

	return isObjectNearbyTargetHorizontal && isObjectNearbyTargetVertical
}
