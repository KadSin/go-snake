package interaction

import (
	"kadsin/shoot-run/game/assets"
	"strings"
	"unicode/utf8"

	term "github.com/nsf/termbox-go"
)

const (
	ALIGNMENT_LEFT = iota
	ALIGNMENT_RIGHT
	ALIGNMENT_CENTER
)

type Content struct {
	Text      string
	Position  assets.Coordinate
	Color     term.Attribute
	Alignment int
}

func (content *Content) Print() {
	texts := strings.Split(content.Text, "\n")

	content.Position.Y -= len(texts) / 2
	for _, t := range texts {
		content.println(t)

		content.Position.Y++
	}
}

func (content Content) println(text string) {
	x := content.Position.X

	switch content.Alignment {
	case ALIGNMENT_RIGHT:
		x -= len(text)
	case ALIGNMENT_CENTER:
		x -= len(text) / 2
	}

	chars := []rune(text)

	for _, char := range chars {
		term.SetCell(
			x, content.Position.Y,
			rune(char),
			content.Color,
			term.GetCell(x, content.Position.Y).Bg,
		)

		x += utf8.RuneLen(char)
	}
}
