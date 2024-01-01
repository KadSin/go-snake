package assets

import term "github.com/nsf/termbox-go"

const (
	SPEED_SHOOTER          = 40
	SPEED_BULLET           = 6
	SPEED_MIN_ENEMY        = 80
	SPEED_MAX_ENEMY        = 125
	SPEED_ENEMY_GENERATOR  = 1000
	SPEED_BLOCKS_GENERATOR = 10000
)

const (
	COLOR_WALLS      = term.ColorGreen
	COLOR_ENEMIES    = term.ColorRed
	COLOR_BULLETS    = term.ColorLightGray
	COLOR_BACKGROUND = term.ColorBlack
)

const (
	IMPACT_SHOOT_ON_ENEMY_GENERATING = 25
)

const (
	KILL_TIMES_TO_SHOW_ENEMY_INCREASING_STORY = 3
)
