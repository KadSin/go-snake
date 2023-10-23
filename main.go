package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

type Animal struct {
	X     int
	Y     int
	Shape rune
}

var snake = Animal{X: 0, Y: 0, Shape: '●'}

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

Infinite:
	for {
		term.Clear(term.ColorDefault, term.ColorDefault)
		term.SetChar(snake.X, snake.Y, snake.Shape)
		term.Sync()
		fmt.Print(" ", snake.X, snake.Y)

		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				snake.moveLeft()
			case term.KeyArrowRight:
				snake.moveRight()
			case term.KeyArrowUp:
				snake.moveUp()
			case term.KeyArrowDown:
				snake.moveDown()
			case term.KeyCtrlC:
				break Infinite
			}
		}
	}
}

func (animal *Animal) moveLeft() {
	if animal.isNotOnBoundry('l') {
		animal.X -= 1
	}
}

func (animal *Animal) moveRight() {
	if animal.isNotOnBoundry('r') {
		animal.X += 1
	}
}

func (animal *Animal) moveUp() {
	if animal.isNotOnBoundry('u') {
		animal.Y -= 1
	}
}

func (animal *Animal) moveDown() {
	if animal.isNotOnBoundry('d') {
		animal.Y += 1
	}
}

func (animal *Animal) isNotOnBoundry(direction rune) bool {
	var width, height = term.Size()

	switch direction {
	case 'l':
		return animal.X > 0
	case 'r':
		return animal.X < width
	case 'u':
		return animal.Y > 0
	case 'd':
		return animal.Y < height
	}

	panic("Bad direction")
}
