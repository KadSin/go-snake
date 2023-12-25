package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
	"math/rand"
	"time"
)

func (game *Game) update() {
	ticker := time.NewTicker(time.Millisecond)

	for range ticker.C {
		if game.Exited {
			break
		}

		game.moveShooter()

		game.generateEnemy()
		game.moveEnemies()

		game.moveBullets()

		game.render()
	}
}

func (game *Game) moveShooter() {
	if game.isTimeToMoveShooter() {
		game.Shooter.Person.UpdateLocation(1)
	}
}

func (game *Game) generateEnemy() {
	if game.isTimeToGenerateEnemy() {
		x := randomElement(game.Screen.Start.X, game.Screen.End.X)
		y := randomElement(game.Screen.Start.Y, game.Screen.End.Y)

		if rand.Float32() > 0.5 {
			x = randomNumberBetween(game.Screen.Start.X, game.Screen.End.X)
		} else {
			y = randomNumberBetween(game.Screen.Start.Y, game.Screen.End.Y)
		}

		enemy := entities.Enemy{
			Person: entities.Object{
				Shape:    '#',
				Location: assets.Coordinate{X: x, Y: y},
				Screen:   game.Screen,
				Color:    assets.COLOR_ENEMIES,
			},
			Target: &game.Shooter.Person,
			Speed:  randomNumberBetween(assets.SPEED_MIN_ENEMY, assets.SPEED_MAX_ENEMY),
		}

		game.Enemies = append(game.Enemies, &enemy)
		game.LastTimeActions.Enemies[&enemy] = 0
	}
}

func randomElement(first int, second int) int {
	if rand.Float32() > 0.5 {
		return first
	} else {
		return second
	}
}

func randomNumberBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}

func (game *Game) moveEnemies() {
	for _, e := range game.Enemies {
		if game.isTimeToMoveEnemy(e) {
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
}

func (game *Game) moveBullets() {
	if game.isTimeToMoveBullet() {
		for _, b := range game.Shooter.Bullets {
			game.Shooter.GoShot(b)

			if game.anEnemyHitBy(b) {
				game.KilledEnemiesCount++

				game.Shooter.RemoveBullet(b)
			}
		}
	}
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

func (game *Game) isTimeToGenerateEnemy() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.EnemyGenerator+int64(assets.SPEED_ENEMY_GENERATOR) {
		game.LastTimeActions.EnemyGenerator = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) isTimeToMoveShooter() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Shooter+int64(game.Shooter.Speed) {
		game.LastTimeActions.Shooter = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) isTimeToMoveEnemy(enemy *entities.Enemy) bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Enemies[enemy]+int64(enemy.Speed) {
		game.LastTimeActions.Enemies[enemy] = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) isTimeToMoveBullet() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Bullets+int64(assets.SPEED_BULLET) {
		game.LastTimeActions.Bullets = time.Now().UnixMilli()
		return true
	}

	return false
}
