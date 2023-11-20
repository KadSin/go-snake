package entities

import (
	"time"
)

type Enemy struct {
	Person Object
	Target *Object
}

func (enemy *Enemy) GoKill(speed int) {
	ticker := time.NewTicker(time.Second / time.Duration(speed))

	for range ticker.C {
		enemy.walk()
	}
}

func (enemy *Enemy) walk() {
	if enemy.Target.X > enemy.Person.X {
		enemy.Person.MoveRight()
	} else if enemy.Target.X < enemy.Person.X {
		enemy.Person.MoveLeft()
	}
	enemy.Person.UpdateLocation(1)

	if enemy.Target.Y > enemy.Person.Y {
		enemy.Person.MoveDown()
	} else if enemy.Target.Y < enemy.Person.Y {
		enemy.Person.MoveUp()
	}
	enemy.Person.UpdateLocation(1)
}
