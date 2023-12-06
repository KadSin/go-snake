package main

import (
	"kadsin/shoot-run/game"
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"

	term "github.com/nsf/termbox-go"
)

func main() {
	var g game.Game

	term.Init()
	term.HideCursor()
	defer term.Close()

	width, height := term.Size()
	g = game.Game{
		Screen: assets.Screen{
			Start: assets.Coordinate{X: 1, Y: 1},
			End:   assets.Coordinate{X: width - 1, Y: height - 1},
		},
	}
	g.Shooter = entities.Shooter{
		Person: entities.Object{
			Shape:     '●',
			Direction: entities.DIRECTION_RIGHT,
			Color:     term.ColorYellow,
			Screen:    g.Screen,
		},
	}

	g.Start()
}
