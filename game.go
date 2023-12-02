package main

import (
	"kadsin/shoot-run/entities"
	"math/rand"
	"time"

	term "github.com/nsf/termbox-go"
)

type Game struct {
	Screen  entities.Screen
	Exited  bool
	Shooter entities.Shooter
	Enemies []*entities.Enemy
}

func (game *Game) Start() {
	game.Shooter.Person.Location = entities.Coordinate{
		X: game.Screen.End.X / 2,
		Y: game.Screen.End.Y / 2,
	}

	go game.listenToKeyboard()

	go game.Shooter.Run(24)
	go game.Shooter.ListenToBullets(150)

	go game.generateEnemies()
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
		}

		go enemy.GoKill(randomNumberBetween(8, 12), func() {
			game.Exited = true
		})

		game.Enemies = append(game.Enemies, &enemy)
	}
}

func randomNumberBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}
