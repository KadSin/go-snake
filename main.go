package main

import (
	"fmt"

	"atomicgo.dev/cursor"
	"github.com/inancgumus/screen"
	term "github.com/nsf/termbox-go"
)

var animal = "|"
var x, y = 0, 0

func main() {
	var width, height = screen.Size()
	x = width / 2
	y = height / 2

	changePosition(x, y)
}

func changePosition(x int, y int) {
	term.Init()
	term.HideCursor()

	defer term.Close()

	for {
		cursor.Move(x, y)

		fmt.Print(animal)

		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				term.Sync()
				fmt.Print(x, y)

				if x > 0 {
					x -= 1
				} else {
					x = 0
				}
			case term.KeyCtrlC:
				panic("bye")
			default:
				term.Flush()
			}
		}
	}
}
