package main

import (
	"kadsin/shoot-run/entities"
	"time"

	term "github.com/nsf/termbox-go"
)

var shooter = entities.Shooter{
	Person: entities.Object{Shape: '●', Direction: entities.DIRECTION_RIGHT, Color: term.ColorYellow},
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

	go shooter.Run(24)

	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		if exit {
			break
		}

		term.Clear(term.ColorDefault, term.ColorDefault)

		PrintObject(shooter.Person)

		for _, bullet := range shooter.Bullets {
			PrintObject(*bullet)
		}

		term.Flush()
	}
}

func listenToKeyboard() {
	for {
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
				go shooter.Shoot(150)
			case term.KeyCtrlC:
				exit = true
			}
		}
	}
}

func PrintObject(object entities.Object) {
	term.SetCell(object.X, object.Y, object.Shape, object.Color, term.ColorDefault)
}
