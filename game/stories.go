package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/interaction"

	term "github.com/nsf/termbox-go"
)

func (game Game) storyReady() interaction.Story {
	return interaction.Story{
		Content: interaction.Content{
			Text:      "🤵: You will gonna kill zombies with the [SPACE] key\n\nIf you wanna suicide then press the [Ctrl]+[C]",
			Position:  assets.Coordinate{X: game.Screen.End.X / 2, Y: game.Screen.End.Y / 2},
			Alignment: interaction.ALIGNMENT_CENTER,
			Color:     term.ColorWhite,
		},
		Background: term.ColorBlack,
		PassMethod: interaction.PASS_BY_KEY,
		KeyToPass:  term.KeySpace,
	}
}

func (game Game) storyGameOver() interaction.Story {
	return interaction.Story{
		Content: interaction.Content{
			Text:      "🤦: Huh, You lost everything...\n\nYou are gonna be a zombie 😏",
			Position:  assets.Coordinate{X: game.Screen.End.X / 2, Y: game.Screen.End.Y / 2},
			Alignment: interaction.ALIGNMENT_CENTER,
			Color:     term.ColorRed,
		},
		Background:    term.ColorBlack,
		PassMethod:    interaction.PASS_BY_TTL,
		SecondsToLive: 3,
	}
}
