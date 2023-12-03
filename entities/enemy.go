package entities

type Enemy struct {
	Person Object
	Target *Object
	Speed  int
	OnKill func()
}

func (enemy *Enemy) Walk() {
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
