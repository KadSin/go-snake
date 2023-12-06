package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"

	term "github.com/nsf/termbox-go"
)

type Game struct {
	Screen          assets.Screen
	Exited          bool
	Shooter         entities.Shooter
	Enemies         []*entities.Enemy
	LastTimeActions LastActionAt
}

type LastActionAt struct {
	Enemies        map[*entities.Enemy]int64
	EnemyGenerator int64
	Shooter        int64
	Bullets        int64
}

func (game *Game) Start() {
	game.LastTimeActions.Enemies = make(map[*entities.Enemy]int64)

	game.Shooter.Speed = assets.SPEED_SHOOTER
	game.Shooter.Person.Location = assets.Coordinate{
		X: game.Screen.End.X / 2,
		Y: game.Screen.End.Y / 2,
	}

	go game.listenToKeyboard()

	game.update()
}

func (game *Game) listenToKeyboard() {
	for {
		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				game.Shooter.Person.MoveLeft()
			case term.KeyArrowRight:
				game.Shooter.Person.MoveRight()
			case term.KeyArrowUp:
				game.Shooter.Person.MoveUp()
			case term.KeyArrowDown:
				game.Shooter.Person.MoveDown()
			case term.KeySpace:
				go game.Shooter.Shoot()
			case term.KeyCtrlC:
				game.Exited = true
			}
		}
	}
}

func (game *Game) render() {
	printObject(game.Shooter.Person)

	for _, bullet := range game.Shooter.Bullets {
		printObject(*bullet)
	}

	for _, enemy := range game.Enemies {
		printObject(enemy.Person)
	}

	game.drawWalls()

	term.Flush()
	term.Clear(term.ColorDefault, term.ColorDefault)
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
