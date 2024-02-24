package entities

import (
	"kadsin/shoot-run/game/assets"
	"slices"
)

type Shooter struct {
	Person  Object
	Speed   int
	Bullets []*Object
	Blood   int
}

func (shooter *Shooter) Shoot() {
	bullet := &Object{
		Shape:     assets.SHAPE_BULLET,
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
	shooter.Bullets = slices.DeleteFunc[[]*Object, *Object](
		shooter.Bullets,
		func(b *Object) bool { return b == bullet },
	)
}

func (shooter *Shooter) State() string {
	stateCount := len(assets.SHAPE_SHOOTER_STATES)
	if shooter.Blood > stateCount-1 {
		return assets.SHAPE_SHOOTER_STATES[stateCount-1]
	}

	return assets.SHAPE_SHOOTER_STATES[shooter.Blood]
}
