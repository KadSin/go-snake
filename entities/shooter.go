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

		term.SetCell(shooter.Bullet.X, shooter.Bullet.Y, shooter.Bullet.Shape, shooter.Bullet.Color, term.ColorDefault)
		term.Sync()
	}
}
