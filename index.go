package main

import (
	"kadsin/shoot-run/entities"
	"time"

	term "github.com/nsf/termbox-go"
)

var game Game

func main() {
	term.Init()
	term.HideCursor()
	defer term.Close()

	width, height := term.Size()
	game = Game{
		Screen: entities.Screen{
			Start: entities.Coordinate{X: 1, Y: 1},
			End:   entities.Coordinate{X: width - 1, Y: height - 1},
		},
	}
	game.Shooter = entities.Shooter{
		Person: entities.Object{
			Shape:     '●',
			Direction: entities.DIRECTION_RIGHT,
			Color:     term.ColorYellow,
			Screen:    game.Screen,
		},
	}

	game.Start()
	render()
}

func render() {

	ticker := time.NewTicker(time.Millisecond)

	for range ticker.C {
		if game.Exited {
			break
		}

		printObject(game.Shooter.Person)

		for _, bullet := range game.Shooter.Bullets {
			printObject(*bullet)
		}

		for _, enemy := range game.Enemies {
			printObject(enemy.Person)
		}

		drawWalls()

		term.Flush()
		term.Clear(term.ColorDefault, term.ColorDefault)
	}
}

func drawWalls() {
	for x := game.Screen.Start.X - 1; x < game.Screen.End.X+1; x++ {
		term.SetCell(x, game.Screen.Start.Y-1, '█', term.ColorGreen, term.ColorDefault)
		term.SetCell(x, game.Screen.End.Y, '█', term.ColorGreen, term.ColorDefault)
	}

	for y := game.Screen.Start.Y - 1; y < game.Screen.End.Y+1; y++ {
		term.SetCell(game.Screen.Start.X-1, y, '█', term.ColorGreen, term.ColorDefault)
		term.SetCell(game.Screen.End.X, y, '█', term.ColorGreen, term.ColorDefault)
	}
}

func printObject(object entities.Object) {
	term.SetCell(object.Location.X, object.Location.Y, object.Shape, object.Color, term.ColorDefault)
}
