package entities

import (
	"kadsin/shoot-run/game/helpers"
)

type Enemy struct {
	Person Object
	Target *Object
	Speed  int
	OnKill func()
}

func (enemy *Enemy) LookAtTarget() {
	canChaseTwoDirections := helpers.RandomBoolean()
	canChaseVertical := helpers.RandomBoolean()

	if canChaseVertical || canChaseTwoDirections {
		enemy.setVerticalDirection()
	}

	if !canChaseVertical || canChaseTwoDirections {
		enemy.setHorizontalDirection()
	}
}

func (enemy *Enemy) setVerticalDirection() {
	if enemy.Target.Location.Y > enemy.Person.Location.Y {
		enemy.Person.MoveDown()
	} else if enemy.Target.Location.Y < enemy.Person.Location.Y {
		enemy.Person.MoveUp()
	}
}

func (enemy *Enemy) setHorizontalDirection() {
	if enemy.Target.Location.X > enemy.Person.Location.X {
		enemy.Person.AdditionalMoveRight()
	} else if enemy.Target.Location.X < enemy.Person.Location.X {
		enemy.Person.AdditionalMoveLeft()
	}
}
