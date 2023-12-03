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

const (
	SPEED_SHOOTER   = 40
	SPEED_BULLET    = 6
	SPEED_MIN_ENEMY = 80
	SPEED_MAX_ENEMY = 125
)

func (game *Game) Start() {
	game.LastTimeMovement.Enemies = make(map[*entities.Enemy]int64)

	game.Shooter.Speed = SPEED_SHOOTER
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
			Speed:  randomNumberBetween(SPEED_MIN_ENEMY, SPEED_MAX_ENEMY),
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
		game.moveShooter(t)
		game.moveEnemies(t)
		game.moveBullets(t)
	}
}

func (game *Game) moveShooter(t time.Time) {
	if t.UnixMilli() > game.LastTimeMovement.Shooter+int64(game.Shooter.Speed) {
		game.Shooter.Person.UpdateLocation(1)
		game.LastTimeMovement.Shooter = t.UnixMilli()
	}
}

func (game *Game) moveEnemies(t time.Time) {
	for _, e := range game.Enemies {
		if t.UnixMilli() > game.LastTimeMovement.Enemies[e]+int64(e.Speed) {
			e.Walk()

			if e.Person.DoesHit(*e.Target) {
				game.Exited = true
			}

			game.LastTimeMovement.Enemies[e] = t.UnixMilli()
		}
	}
}

func (game *Game) moveBullets(t time.Time) {
	if t.UnixMilli() > game.LastTimeMovement.Bullets+SPEED_BULLET {
		for _, b := range game.Shooter.Bullets {
			game.Shooter.GoShot(b)

			for _, e := range game.Enemies {
				if b.DoesHit(e.Person) {
					game.RemoveEnemy(e)
					game.Shooter.RemoveBullet(b)
				}
			}
		}

		game.LastTimeMovement.Bullets = t.UnixMilli()
	}
}

func (game *Game) RemoveEnemy(enemy *entities.Enemy) {
	for id, e := range game.Enemies {
		if e == enemy {
			game.Enemies[id] = nil

			if id == 0 {
				game.Enemies = game.Enemies[id+1:]
			} else if id == len(game.Enemies)-1 {
				game.Enemies = game.Enemies[:id-1]
			} else {
				game.Enemies = append(game.Enemies[id-1:], game.Enemies[:id]...)
			}

			break
		}
	}
}
