package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
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
