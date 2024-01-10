package game

import (
	"fmt"
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/interaction"
	"time"

	term "github.com/nsf/termbox-go"
)

func (game Game) showStoryReady() {
	game.storyHelpToShoot().Show()
	game.storyHelpToExit().Show()
}

func (game Game) storyHelpToShoot() interaction.Story {
	return game.storyByTtl(
		"🤵: You will gonna kill zombies with the [SPACE] key",
		2,
		term.ColorLightRed,
	)
}

func (game Game) storyHelpToExit() interaction.Story {
	return game.storyByTtl(
		"🤵: If you wanna suicide then press the [Ctrl]+[C]",
		2,
		term.ColorLightYellow,
	)
}

func (game Game) storyGameOver() interaction.Story {
	return game.storyByTtl(
		"🤦: Huh, You lost everything...\n\nYou are gonna be a zombie 😏",
		3,
		term.ColorRed,
	)
}

func (game Game) storyByTtl(text string, seconds int, color term.Attribute) interaction.Story {
	return interaction.Story{
		Content: interaction.Content{
			Text:      text,
			Position:  assets.Coordinate{X: game.Screen.End.X / 2, Y: game.Screen.End.Y / 2},
			Alignment: interaction.ALIGNMENT_CENTER,
			Color:     color,
		},
		Background:    term.ColorBlack,
		PassMethod:    interaction.PASS_BY_TTL,
		SecondsToLive: seconds,
	}
}

func (game Game) storyHelpAboutSpeedOfZombies() interaction.Story {
	return game.storyByKey(
		"🤵: If you kill more then more zombies will know you are here\nMind the awareness percentage below the screen",
		term.ColorLightRed,
	)
}

func (game Game) storyShowScore() interaction.Story {
	rectangleCircumference := game.ScreenCircumference()
	score := (time.Now().Unix() - game.StartedAt) * int64(game.KilledEnemiesCount) / int64(rectangleCircumference/10)

	level, _ := game.screenLevel()

	return game.storyByKey(
		fmt.Sprintf(
			"Score: %v\n\nScreen Level: %v\nTotal Killed: %v\nLife time: %v",
			score, level, game.KilledEnemiesCount, game.ScreenTime(),
		),
		term.ColorWhite,
	)
}

func (game Game) storyByKey(text string, color term.Attribute) interaction.Story {
	return interaction.Story{
		Content: interaction.Content{
			Text:      text + "\n\nPress [Enter] to continue",
			Position:  assets.Coordinate{X: game.Screen.End.X / 2, Y: game.Screen.End.Y / 2},
			Alignment: interaction.ALIGNMENT_CENTER,
			Color:     color,
		},
		Background: term.ColorBlack,
		PassMethod: interaction.PASS_BY_KEY,
		KeyToPass:  term.KeyEnter,
	}
}
