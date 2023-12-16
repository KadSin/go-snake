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
