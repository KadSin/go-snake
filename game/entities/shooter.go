package entities

import (
	"kadsin/shoot-run/game/assets"
)

type Shooter struct {
	Person  Object
	Speed   int
	Bullets []*Object
	Blood   int
}

func (shooter *Shooter) Shoot() {
	bullet := &Object{
		Shape:     '*',
		Direction: shooter.Person.Direction,
		Color:     assets.COLOR_BULLETS,
		Location:  shooter.Person.Location,
		Screen:    shooter.Person.Screen,
	}

	shooter.Bullets = append(shooter.Bullets, bullet)

	bullet.UpdateLocation(2)
}

func (shooter *Shooter) GoShot(bullet *Object) {
	error := bullet.UpdateLocation(1)

	if error != nil {
		shooter.RemoveBullet(bullet)
	}
}

func (shooter *Shooter) RemoveBullet(bullet *Object) {
	for id, b := range shooter.Bullets {
		if b == bullet {
			shooter.Bullets[id] = nil

			if id == 0 {
				shooter.Bullets = shooter.Bullets[id+1:]
			} else if id == len(shooter.Bullets)-1 {
				shooter.Bullets = shooter.Bullets[:id-1]
			} else {
				shooter.Bullets = append(shooter.Bullets[:id], shooter.Bullets[id+1:]...)
			}

			break
		}
	}
}
