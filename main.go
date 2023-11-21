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

	game.Width, game.Height = term.Size()
	game.Shooter = entities.Shooter{
		Person: entities.Object{
			Shape:     '●',
			Direction: entities.DIRECTION_RIGHT,
			Color:     term.ColorYellow,
			MaxX:      game.Width,
			MaxY:      game.Height,
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

		term.Flush()
		term.Clear(term.ColorDefault, term.ColorDefault)
	}
}

func printObject(object entities.Object) {
	term.SetCell(object.X, object.Y, object.Shape, object.Color, term.ColorDefault)
}
