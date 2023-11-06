package main

import (
	"kadsin/shoot-run/entities"
	"time"

	term "github.com/nsf/termbox-go"
)

var shooter = entities.Shooter{
	Person: entities.Object{Shape: '●', Direction: entities.DIRECTION_RIGHT, Color: term.ColorYellow},
	Bullet: entities.Object{Shape: '*', Color: term.ColorLightGray},
}

var exit = false

func main() {
	term.Init()
	term.HideCursor()
	defer term.Close()

	startGame()
}

func startGame() {
	var width, height = term.Size()
	shooter.Person.X = width / 2
	shooter.Person.Y = height / 2

	go listenToKeyboard()

	ticker := time.NewTicker(time.Second / 24)

	for range ticker.C {
		if exit {
			break
		}

		shooter.Walk()
	}
}

func listenToKeyboard() {
	for {
		if exit {
			break
		}

		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				shooter.Person.MoveLeft()
			case term.KeyArrowRight:
				shooter.Person.MoveRight()
			case term.KeyArrowUp:
				shooter.Person.MoveUp()
			case term.KeyArrowDown:
				shooter.Person.MoveDown()
			case term.KeySpace:
				if !shooter.IsShooting {
					go shooter.Shoot(150)
				}
			case term.KeyCtrlC:
				exit = true
			}
		}
	}
}
