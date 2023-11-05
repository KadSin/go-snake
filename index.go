package main

import (
	"kadsin/snake/entities"
	"time"

	term "github.com/nsf/termbox-go"
)

var snake = entities.Animal{X: 0, Y: 0, Shape: '●', Direction: entities.DIRECTION_RIGHT}
var exit = false

func main() {
	term.Init()
	term.HideCursor()
	defer term.Close()

	startGame()
}

func startGame() {
	var width, height = term.Size()
	snake.X = width / 2
	snake.Y = height / 2

	ticker := time.NewTicker(time.Second / 24) // == 24 FPS

	go listenToKeyboard()

	for range ticker.C {
		if exit {
			break
		}

		term.Clear(term.ColorDefault, term.ColorDefault)
		term.SetChar(snake.X, snake.Y, snake.Shape)
		term.Sync()

		snake.UpdateLocation()
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
				snake.MoveLeft()
			case term.KeyArrowRight:
				snake.MoveRight()
			case term.KeyArrowUp:
				snake.MoveUp()
			case term.KeyArrowDown:
				snake.MoveDown()
			case term.KeyCtrlC:
				exit = true
			}
		}
	}
}
