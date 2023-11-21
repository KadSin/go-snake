package entities

import (
	"time"
)

type Enemy struct {
	Person Object
	Target *Object
}

func (enemy *Enemy) GoKill(speed int, onKill func()) {
	ticker := time.NewTicker(time.Second / time.Duration(speed))

	for range ticker.C {
		enemy.walk()

		if enemy.didKill() {
			onKill()
		}
	}
}

func (enemy *Enemy) walk() {
	if enemy.Target.Location.X > enemy.Person.Location.X {
		enemy.Person.MoveRight()
	} else if enemy.Target.Location.X < enemy.Person.Location.X {
		enemy.Person.MoveLeft()
	}
	enemy.Person.UpdateLocation(1)

	if enemy.Target.Location.Y > enemy.Person.Location.Y {
		enemy.Person.MoveDown()
	} else if enemy.Target.Location.Y < enemy.Person.Location.Y {
		enemy.Person.MoveUp()
	}
	enemy.Person.UpdateLocation(1)
}

func (enemy *Enemy) didKill() bool {
	if enemy.Person.Location.X == enemy.Target.Location.X && enemy.Person.Location.Y == enemy.Target.Location.Y {
		return true
	}

	return false
}
