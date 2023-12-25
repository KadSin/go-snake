package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
	"kadsin/shoot-run/game/interaction"
	"strconv"
	"strings"

	term "github.com/nsf/termbox-go"
)

func (game *Game) render() {
	game.drawEntities()

	game.drawTopBar()
	game.drawKilledEnemiesCount()
	game.drawBlood()

	game.drawWalls()

	term.Flush()
	term.Clear(term.ColorDefault, term.ColorDefault)
}

func (game *Game) drawEntities() {
	printObject(game.Shooter.Person)

	for _, bullet := range game.Shooter.Bullets {
		printObject(*bullet)
	}

	for _, enemy := range game.Enemies {
		printObject(enemy.Person)
	}
}

func printObject(object entities.Object) {
	term.SetCell(object.Location.X, object.Location.Y, object.Shape, object.Color, term.ColorDefault)
}

func (game *Game) drawTopBar() {
	for i := game.Screen.Start.X - 1; i < game.Screen.End.X+1; i++ {
		term.SetBg(i, game.Screen.Start.Y-2, term.ColorBlack)
	}
}

func (game *Game) drawKilledEnemiesCount() {
	killedEnemiesCount := "💀" + strconv.FormatInt(int64(game.KilledEnemiesCount), 10)

	content := interaction.Content{
		Position:  assets.Coordinate{X: game.Screen.End.X, Y: game.Screen.Start.Y - 2},
		Text:      killedEnemiesCount,
		Alignment: interaction.ALIGNMENT_RIGHT,
		Color:     term.ColorWhite,
	}
	content.Print()
}

func (game *Game) drawBlood() {
	content := interaction.Content{
		Position: assets.Coordinate{X: game.Screen.Start.X, Y: game.Screen.Start.Y - 2},
		Text:     game.state() + strings.Repeat("♥", game.Shooter.Blood),
		Color:    term.ColorRed,
	}
	content.Print()
}

func (game *Game) state() string {
	states := []string{"😖", "😨", "😐", "😀", "😄", "😁"}

	if game.Shooter.Blood > len(states) {
		return "😇"
	}

	return states[game.Shooter.Blood]

}

func (game *Game) drawWalls() {
	for x := game.Screen.Start.X - 1; x < game.Screen.End.X+1; x++ {
		term.SetCell(x, game.Screen.Start.Y-1, '█', assets.COLOR_WALLS, term.ColorDefault)
		term.SetCell(x, game.Screen.End.Y, '█', assets.COLOR_WALLS, term.ColorDefault)
	}

	for y := game.Screen.Start.Y - 1; y < game.Screen.End.Y+1; y++ {
		term.SetCell(game.Screen.Start.X-1, y, '█', assets.COLOR_WALLS, term.ColorDefault)
		term.SetCell(game.Screen.End.X, y, '█', assets.COLOR_WALLS, term.ColorDefault)
	}
}
