package main

import (
	"fmt"
	"kadsin/snake/entities"
	"time"

	term "github.com/nsf/termbox-go"
)

var snake = entities.Animal{X: 0, Y: 0, Shape: '●', Direction: entities.DIRECTION_RIGHT}

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

Infinite:
	for range ticker.C {
		term.Clear(term.ColorDefault, term.ColorDefault)
		term.SetChar(snake.X, snake.Y, snake.Shape)
		term.Sync()
		fmt.Print(" ", snake.X, snake.Y)

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
				break Infinite
			}
		}

		snake.UpdateLocation()
	}
}
