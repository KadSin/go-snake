package entities

import term "github.com/nsf/termbox-go"

type Animal struct {
	X     int
	Y     int
	Shape rune
}

func (animal *Animal) MoveLeft() {
	if animal.isNotOnBoundry('l') {
		animal.X -= 1
	}
}

func (animal *Animal) MoveRight() {
	if animal.isNotOnBoundry('r') {
		animal.X += 1
	}
}

func (animal *Animal) MoveUp() {
	if animal.isNotOnBoundry('u') {
		animal.Y -= 1
	}
}

func (animal *Animal) MoveDown() {
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
		return animal.X < width-1
	case 'u':
		return animal.Y > 0
	case 'd':
		return animal.Y < height-1
	}

	panic("Bad direction")
}
