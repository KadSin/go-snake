package game

import (
	"errors"
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
	"kadsin/shoot-run/game/helpers"
	"slices"
	"time"
)

func (game *Game) EventCollisionShooterByEnemy(enemy *entities.Enemy) {
	if game.Shooter.Blood > 0 {
		game.Shooter.Blood--

		game.removeEnemy(enemy)
	} else {
		game.storyGameOver().Show()

		game.Exited = true
	}
}

func (game *Game) EventCollisionEnemyByBullet(enemy *entities.Enemy, bullet *entities.Object) {
	if game.KilledEnemiesCount == assets.KILL_TIMES_TO_SHOW_ENEMY_INCREASING_STORY {
		game.storyHelpAboutSpeedOfZombies().Show()
	}

	game.KilledEnemiesCount++
	game.removeEnemy(enemy)
	game.LastTimeActions.Kill = time.Now().UnixMilli()

	game.Shooter.RemoveBullet(bullet)

	game.increaseEnemyGeneratorSpeed()
}

func (game *Game) increaseEnemyGeneratorSpeed() {
	nextSpeed := game.SpeedEnemyGenerator - assets.IMPACT_SHOOT_ON_ENEMY_GENERATING*10

	if nextSpeed >= assets.SPEED_MIN_ENEMY_GENERATOR {
		game.SpeedEnemyGenerator = nextSpeed
	}
}

func (game *Game) removeEnemy(enemy *entities.Enemy) {
	game.Enemies = slices.DeleteFunc[[]*entities.Enemy, *entities.Enemy](
		game.Enemies,
		func(e *entities.Enemy) bool { return e == enemy },
	)
}

func (game *Game) EventCollisionBlockByBullet(block *entities.Object, bullet *entities.Object) {
	game.Shooter.RemoveBullet(bullet)
}

func (game *Game) EventCollisionBlockByShooter(block *entities.Object) error {
	return errors.New("Shooter should stop")
}

func (game *Game) EventCollisionBlockByEnemy(block *entities.Object, enemy *entities.Enemy) error {
	return errors.New("Enemy should stop")
}

func (game *Game) EventCollisionPortalByShooter() {
	game.generateBlocks()

	game.Portal.Location = helpers.RandomCoordinate(game.Screen, assets.Coordinate{X: 1, Y: 1})

	game.Enemies = make([]*entities.Enemy, 0)
}

func (game *Game) generateBlocks() {
	game.Blocks = []entities.Object{}

	count := helpers.RandomNumberBetween(3, 5)

	for i := 0; i < count; i++ {
		size := helpers.RandomNumberBetween(3, 15)
		location := helpers.RandomCoordinate(game.Screen, assets.Coordinate{X: 2, Y: 2})

		for j := 0; j < size; j++ {
			isHorizontal := helpers.RandomBoolean()

			shape := '█'
			if isHorizontal {
				shape = '▀'
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
