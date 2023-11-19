package entities

import (
	term "github.com/nsf/termbox-go"
)

func printObject(object Object) {
	term.SetCell(object.X, object.Y, object.Shape, object.Color, term.ColorDefault)

	term.Sync()
}
