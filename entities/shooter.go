package entities

import (
	"time"

	term "github.com/nsf/termbox-go"
)

type Shooter struct {
	Person     Object
	Bullet     Object
	IsShooting bool
}

func (shooter *Shooter) Shoot(speed int) {
	shooter.IsShooting = true

	shooter.Bullet.Direction = shooter.Person.Direction
	shooter.Bullet.X = shooter.Person.X
	shooter.Bullet.Y = shooter.Person.Y
	shooter.Bullet.UpdateLocation(2)

	ticker := time.NewTicker(time.Second / time.Duration(speed))

	for range ticker.C {
		term.SetChar(shooter.Bullet.X, shooter.Bullet.Y, ' ')

		error := shooter.Bullet.UpdateLocation(1)
		if error != nil {
			shooter.IsShooting = false
			break
		}

		printObject(shooter.Bullet)
	}
}

func (shooter *Shooter) Walk() {
	term.SetChar(shooter.Person.X, shooter.Person.Y, ' ')

	shooter.Person.UpdateLocation(1)
	printObject(shooter.Person)
}

func printObject(object Object) {
	term.SetCell(object.X, object.Y, object.Shape, object.Color, term.ColorDefault)

	term.Sync()
}
