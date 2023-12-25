package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"

	term "github.com/nsf/termbox-go"
)

func (game *Game) render() {
	game.drawWalls()
	game.drawEntities()

	game.drawTopBar()
	game.drawBlood()
	game.drawState()

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

func (game *Game) drawTopBar() {
	for i := game.Screen.Start.X - 1; i < game.Screen.End.X+1; i++ {
		term.SetBg(i, game.Screen.Start.Y-2, term.ColorBlack)
		term.SetFg(i, game.Screen.Start.Y-2, term.ColorRed)
	}
}

func (game *Game) drawBlood() {
	for i := 0; i < game.Shooter.Blood*2; i += 2 {
		term.SetChar(i+4, 0, '♥')
	}
}

func (game *Game) drawState() {
	states := []rune{'😖', '😨', '😐', '😀', '😄', '😁'}
	state := states[game.Shooter.Blood]

	if game.Shooter.Blood > len(states) {
		state = '😇'
	}

	term.SetChar(1, 0, state)
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

func printObject(object entities.Object) {
	term.SetCell(object.Location.X, object.Location.Y, object.Shape, object.Color, term.ColorDefault)
}
