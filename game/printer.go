package game

import (
	"kadsin/shoot-run/game/assets"
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
	switch content.Alignment {
	case ALIGNMENT_RIGHT:
		content.Position.X -= len(content.Text)
	case ALIGNMENT_CENTER:
		content.Position.X -= len(content.Text) / 2
	}

	for _, char := range []rune(content.Text) {
		term.SetCell(
			content.Position.X, content.Position.Y,
			rune(char),
			content.Color,
			term.GetCell(content.Position.X, content.Position.Y).Bg,
		)

		content.Position.X += utf8.RuneLen(char)
	}
}
