package main

import (
	"kadsin/shoot-run/entities"
	"math/rand"
	"time"

	term "github.com/nsf/termbox-go"
)

type Game struct {
	Screen           entities.Screen
	Exited           bool
	Shooter          entities.Shooter
	Enemies          []*entities.Enemy
	LastTimeMovement EntityMovementLastTime
}

type EntityMovementLastTime struct {
	Enemies map[*entities.Enemy]int64
	Shooter int64
	Bullets int64
}

func (game *Game) Start() {
	game.LastTimeMovement.Enemies = make(map[*entities.Enemy]int64)

	game.Shooter.Speed = 40
	game.Shooter.Person.Location = entities.Coordinate{
		X: game.Screen.End.X / 2,
		Y: game.Screen.End.Y / 2,
	}

	go game.listenToKeyboard()

	go game.generateEnemies()

	go game.updateLocations()
}

func (game *Game) listenToKeyboard() {
	for {
		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				game.Shooter.Person.MoveLeft()
			case term.KeyArrowRight:
				game.Shooter.Person.MoveRight()
			case term.KeyArrowUp:
				game.Shooter.Person.MoveUp()
			case term.KeyArrowDown:
				game.Shooter.Person.MoveDown()
			case term.KeySpace:
				go game.Shooter.Shoot()
			case term.KeyCtrlC:
				game.Exited = true
			}
		}
	}
}

func (game *Game) generateEnemies() {
	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		x := 0
		if rand.Float32() < 0.5 {
			x = game.Screen.End.X
		}

		y := 0
		if rand.Float32() < 0.5 {
			y = game.Screen.End.Y
		}

		enemy := entities.Enemy{
			Person: entities.Object{
				Shape:    '#',
				Location: entities.Coordinate{X: x, Y: y},
				Screen:   game.Screen,
				Color:    term.ColorRed,
			},
			Target: &game.Shooter.Person,
			Speed:  randomNumberBetween(80, 125),
			OnKill: func() { game.Exited = true },
		}

		game.Enemies = append(game.Enemies, &enemy)
		game.LastTimeMovement.Enemies[&enemy] = 0
	}
}

func randomNumberBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (game *Game) updateLocations() {
	ticker := time.NewTicker(time.Millisecond)

	for t := range ticker.C {
		game.moveEnemies(t)
		game.moveShooter(t)
		game.moveBullets(t)
	}
}

func (game *Game) moveEnemies(t time.Time) {
	for _, e := range game.Enemies {
		if t.UnixMilli() > game.LastTimeMovement.Enemies[e]+int64(e.Speed) {
			e.GoKill()
			game.LastTimeMovement.Enemies[e] = t.UnixMilli()
		}
	}
}

func (game *Game) moveShooter(t time.Time) {
	if t.UnixMilli() > game.LastTimeMovement.Shooter+int64(game.Shooter.Speed) {
		game.Shooter.Person.UpdateLocation(1)
		game.LastTimeMovement.Shooter = t.UnixMilli()
	}
}

func (game *Game) moveBullets(t time.Time) {
	if t.UnixMilli() > game.LastTimeMovement.Bullets+6 {
		game.Shooter.UpdateLocationOfBullets()
		game.LastTimeMovement.Bullets = t.UnixMilli()
	}
}
