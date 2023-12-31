package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
	"kadsin/shoot-run/game/helpers"
	"time"
)

func (game *Game) update() {
	ticker := time.NewTicker(time.Millisecond)

	for range ticker.C {
		if game.Exited {
			game.storyShowScore().Show()

			break
		}

		game.generateBlocks()

		game.moveShooter()

		game.generateEnemy()
		game.moveEnemies()

		game.moveBullets()

		game.render()
	}
}

func (game *Game) generateBlocks() {
	if !game.isTimeToGenerateBlocks() {
		return
	}

	game.Blocks = []entities.Object{}

	count := helpers.RandomNumberBetween(10, 15)

	for i := 0; i < count; i++ {
		size := helpers.RandomNumberBetween(3, 6)
		location := assets.Coordinate{
			X: helpers.RandomNumberBetween(game.Screen.Start.X, game.Screen.End.X),
			Y: helpers.RandomNumberBetween(game.Screen.Start.Y, game.Screen.End.Y),
		}

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

func (game *Game) isTimeToGenerateBlocks() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.BlocksGenerator+assets.SPEED_BLOCKS_GENERATOR {
		game.LastTimeActions.BlocksGenerator = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) moveShooter() {
	if game.isTimeToMoveShooter() {
		game.Shooter.Person.UpdateLocation(1)
	}
}

func (game *Game) isTimeToMoveShooter() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Shooter+int64(game.Shooter.Speed) {
		game.LastTimeActions.Shooter = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) generateEnemy() {
	if !game.isTimeToGenerateEnemy() {
		return
	}

	x := helpers.RandomIntElement(game.Screen.Start.X, game.Screen.End.X)
	y := helpers.RandomIntElement(game.Screen.Start.Y, game.Screen.End.Y)

	if helpers.RandomBoolean() {
		x = helpers.RandomNumberBetween(game.Screen.Start.X, game.Screen.End.X)
	} else {
		y = helpers.RandomNumberBetween(game.Screen.Start.Y, game.Screen.End.Y)
	}

	enemy := entities.Enemy{
		Person: entities.Object{
			Shape:    '#',
			Location: assets.Coordinate{X: x, Y: y},
			Screen:   game.Screen,
			Color:    assets.COLOR_ENEMIES,
		},
		Target: &game.Shooter.Person,
		Speed:  helpers.RandomNumberBetween(assets.SPEED_MIN_ENEMY, assets.SPEED_MAX_ENEMY),
	}

	game.Enemies = append(game.Enemies, &enemy)
	game.LastTimeActions.Enemies[&enemy] = 0
}

func (game *Game) isTimeToGenerateEnemy() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.EnemyGenerator+int64(game.enemyGeneratorSpeed()) {
		game.LastTimeActions.EnemyGenerator = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) enemyGeneratorSpeed() uint {
	lastShootDiff := uint(time.Now().UnixMilli()-game.LastTimeActions.Kill) / 100
	if lastShootDiff > 1000 {
		return 1000
	}

	variant := game.KilledEnemiesCount*assets.IMPACT_SHOOT_ON_ENEMY_GENERATING - lastShootDiff
	if variant > 800 {
		return 200
	}

	speed := assets.SPEED_ENEMY_GENERATOR - variant

	return speed
}

func (game *Game) moveEnemies() {
	for _, e := range game.Enemies {
		if !game.isTimeToMoveEnemy(e) {
			continue
		}

		e.Chase()

		if e.Person.DoesHit(*e.Target) {
			if game.Shooter.Blood > 0 {
				game.Shooter.Blood--

				game.removeEnemy(e)
			} else {
				game.storyGameOver().Show()

				game.Exited = true
			}
		}
	}
}

func (game *Game) isTimeToMoveEnemy(enemy *entities.Enemy) bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Enemies[enemy]+int64(enemy.Speed) {
		game.LastTimeActions.Enemies[enemy] = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) moveBullets() {
	if !game.isTimeToMoveBullet() {
		return
	}

	for _, b := range game.Shooter.Bullets {
		game.Shooter.GoShot(b)

		if game.anEnemyHitBy(b) {
			if game.KilledEnemiesCount == 3 {
				game.storyHelpAboutSpeedOfZombies().Show()
			}
			game.KilledEnemiesCount++

			game.LastTimeActions.Kill = time.Now().UnixMilli()

			game.Shooter.RemoveBullet(b)
		}
	}
}

func (game *Game) isTimeToMoveBullet() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Bullets+int64(assets.SPEED_BULLET) {
		game.LastTimeActions.Bullets = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) anEnemyHitBy(bullet *entities.Object) bool {
	for _, e := range game.Enemies {
		if bullet.DoesHit(e.Person) {
			game.removeEnemy(e)

			return true
		}
	}

	return false
}

func (game *Game) removeEnemy(enemy *entities.Enemy) {
	for id, e := range game.Enemies {
		if e == enemy {
			if id == 0 {
				game.Enemies = game.Enemies[id+1:]
			} else if id == len(game.Enemies)-1 {
				game.Enemies = game.Enemies[:id-1]
			} else {
				game.Enemies = append(game.Enemies[:id], game.Enemies[id+1:]...)
			}

			break
		}
	}
}
