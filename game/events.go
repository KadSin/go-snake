package game

import (
	"kadsin/shoot-run/game/assets"
	"kadsin/shoot-run/game/entities"
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
