package game

import (
	"fmt"
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
	"kadsin/shoot-run/game/helpers"
	"time"

	term "github.com/nsf/termbox-go"
)

type Game struct {
	Screen              assets.Screen
	Exited              bool
	Shooter             entities.Shooter
	Enemies             []*entities.Enemy
	KilledEnemiesCount  int
	SpeedEnemyGenerator int
	Blocks              []entities.Object
	Portal              entities.Object
	LastTimeActions     LastActionAt
	StartedAt           int64
}

type LastActionAt struct {
	Portal                      int64
	PortalDirection             int64
	Enemies                     map[*entities.Enemy]int64
	EnemyGenerator              int64
	IncreaseEnemyGeneratorSpeed int64
	Shooter                     int64
	Bullets                     int64
	Kill                        int64
}

func (game *Game) Start() {
	game.showStoryReady()
	game.StartedAt = time.Now().Unix()

	game.SpeedEnemyGenerator = assets.SPEED_MAX_ENEMY_GENERATOR
	game.LastTimeActions.Enemies = make(map[*entities.Enemy]int64)

	game.Shooter.Speed = assets.SPEED_SHOOTER
	game.Shooter.Person.Location = assets.Coordinate{
		X: game.Screen.End.X / 2,
		Y: game.Screen.End.Y / 2,
	}

	game.generateBlocks()
	game.Portal = entities.Object{
		Shape:    '🌀',
		Screen:   game.Screen,
		Location: helpers.RandomCoordinate(game.Screen, assets.Coordinate{X: 1, Y: 1}),
	}

	go game.listenToKeyboard()

	game.update()
}

func (game *Game) generateBlocks() {
	game.Blocks = []entities.Object{}

	count := helpers.RandomNumberBetween(3, 5)

	for i := 0; i < count; i++ {
		size := helpers.RandomNumberBetween(3, 15)
		location := helpers.RandomCoordinate(game.Screen, assets.Coordinate{X: 2, Y: 2})

		for j := 0; j < size; j++ {
			isHorizontal := helpers.RandomBoolean()

			shape := assets.SHAPE_BLOCK_VERTICAL
			if isHorizontal {
				shape = assets.SHAPE_BLOCK_HORIZONTAL
			}

			block := entities.Object{
				Shape:    shape,
				Location: location,
				Screen:   game.Screen,
				Color:    assets.COLOR_WALLS,
			}

			if isHorizontal {
				location.X++
			} else {
				location.Y++
			}

			game.Blocks = append(game.Blocks, block)
		}
	}
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

func (game Game) ScreenTime() string {
	screenTime := time.Now().Unix() - game.StartedAt

	minutes := screenTime / 60
	seconds := screenTime % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func (game *Game) ScreenCircumference() int {
	return 2*game.Screen.End.X + 2*game.Screen.End.Y
}
