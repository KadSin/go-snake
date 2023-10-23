package main

import (
	"fmt"

	"atomicgo.dev/cursor"
	"github.com/inancgumus/screen"
	term "github.com/nsf/termbox-go"
)

var animal = "|"

type Animal struct {
	X     int
	Y     int
	Shape string
}

var snake = Animal{X: 0, Y: 0, Shape: "●"}

func main() {
	var width, height = screen.Size()
	snake.X = width / 2
	snake.Y = height / 2

	startGame()
}

func startGame() {
	term.Init()
	term.HideCursor()
	defer term.Close()

Infinite:
	for {
		cursor.Move(snake.X, snake.Y)
		fmt.Print(snake.Shape)

		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				snake.moveLeft()
			case term.KeyCtrlC:
				break Infinite
			}

			term.Sync()
			fmt.Print(snake.X, snake.Y)
		}
	}
}

func (animal *Animal) moveLeft() {
	if animal.X > 0 {
		animal.X -= 1
	} else {
		animal.X = 0
	}
}
