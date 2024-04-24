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
			return
		}

		game.moveShooter()

		game.decreaseEnemyGeneratorSpeed()
		game.generateEnemy()
		game.moveEnemies()

		game.moveBullets()

		game.changePortalDirection()
		game.movePortal()

		game.render()
	}
}

func (game *Game) moveShooter() {
	if game.isTimeToMoveShooter() {
		if block := game.isShooterBehindOfBlock(); block != nil {
			event := game.EventCollisionBlockByShooter(block)

			if event != nil {
				return
			}
		}

		if game.Shooter.Person.DoesHit(game.Portal) {
			game.EventCollisionPortalByShooter()
		}

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

func (game *Game) isShooterBehindOfBlock() *entities.Object {
	for _, block := range game.Blocks {
		if game.Shooter.Person.NextStep(1) == block.Location {
			return &block
		}
	}

	return nil
}

func (game *Game) decreaseEnemyGeneratorSpeed() {
	if !game.isTimeToIncreaseEnemyGeneratorSpeed() {
		return
	}

	nextSpeed := game.SpeedEnemyGenerator + assets.IMPACT_SHOOT_ON_ENEMY_GENERATING

	if nextSpeed <= assets.SPEED_MAX_ENEMY_GENERATOR {
		game.SpeedEnemyGenerator = nextSpeed
	}
}

func (game *Game) isTimeToIncreaseEnemyGeneratorSpeed() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.IncreaseEnemyGeneratorSpeed+int64(assets.INTERVAL_ENEMY_GENERATOR_SPEED_INCREASER) {
		game.LastTimeActions.IncreaseEnemyGeneratorSpeed = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) generateEnemy() {
	if !game.isTimeToGenerateEnemy() {
		return
	}

	enemy := entities.Enemy{
		Person: entities.Object{
			Shape:    assets.SHAPE_ENEMY,
			Location: helpers.RandomCoordinateOnBorders(game.Screen),
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
	if time.Now().UnixMilli() > game.LastTimeActions.EnemyGenerator+int64(game.SpeedEnemyGenerator) {
		game.LastTimeActions.EnemyGenerator = time.Now().UnixMilli()
		return true
	}

	return false
}

func (game *Game) moveEnemies() {
	for _, e := range game.Enemies {
		if !game.isTimeToMoveEnemy(e) {
			continue
		}

		e.LookAtTarget()

		if block := game.isEnemyBehindOfBlock(e); block != nil {
			event := game.EventCollisionBlockByEnemy(block, e)

			if event != nil {
				continue
			}
		}

		e.Person.UpdateLocation(1)

		if e.Person.DoesHit(*e.Target) {
			game.EventCollisionShooterByEnemy(e)
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

func (game *Game) isEnemyBehindOfBlock(e *entities.Enemy) *entities.Object {
	for _, block := range game.Blocks {
		if e.Person.NextStep(1) == block.Location {
			return &block
		}
	}

	return nil
}

func (game *Game) moveBullets() {
	if !game.isTimeToMoveBullet() {
		return
	}

	for _, b := range game.Shooter.Bullets {
		if block := game.isBulletBehindOfBlock(b); block != nil {
			game.EventCollisionBlockByBullet(block, b)
		}

		game.Shooter.GoShot(b)

		if enemy := game.anEnemyHitBy(b); enemy != nil {
			game.EventCollisionEnemyByBullet(enemy, b)
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

func (game *Game) isBulletBehindOfBlock(bullet *entities.Object) *entities.Object {
	for _, block := range game.Blocks {
		if bullet.DoesHit(block) {
			return &block
		}
	}

	return nil
}

func (game *Game) anEnemyHitBy(bullet *entities.Object) *entities.Enemy {
	for _, e := range game.Enemies {
		if bullet.DoesHit(e.Person) {
			return e
		}
	}

	return nil
}

func (game *Game) movePortal() {
	if !game.isTimeToMovePortal() {
		return
	}

	game.Portal.UpdateLocation(1)
}

func (game *Game) isTimeToMovePortal() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.Portal+assets.SPEED_PORTAL {
		game.LastTimeActions.Portal = time.Now().UnixMilli()

		return true
	}

	return false
}

func (game *Game) changePortalDirection() {
	if !game.isTimeToChangePortalDirection() {
		return
	}

	game.Portal.Direction = uint8(helpers.RandomNumberBetween(1, 4))
}

func (game *Game) isTimeToChangePortalDirection() bool {
	if time.Now().UnixMilli() > game.LastTimeActions.PortalDirection+assets.INTERVAL_PORTAL_DIRECTION {
		game.LastTimeActions.PortalDirection = time.Now().UnixMilli()

		return true
	}

	return false
}
