package main

import (
	"kadsin/shoot-run/entities"
	"math/rand"
	"time"

	term "github.com/nsf/termbox-go"
)

var shooter = entities.Shooter{
	Person: entities.Object{Shape: '●', Direction: entities.DIRECTION_RIGHT, Color: term.ColorYellow},
}

var enemies []*entities.Enemy

var exit = false

func main() {
	term.Init()
	term.HideCursor()
	defer term.Close()

	startGame()
}

func startGame() {
	var width, height = term.Size()
	shooter.Person.X = width / 2
	shooter.Person.Y = height / 2

	go listenToKeyboard()

	go shooter.Run(24)

	go GenerateEnemies()

	ticker := time.NewTicker(time.Millisecond)
	for range ticker.C {
		if exit {
			break
		}

		term.Clear(term.ColorDefault, term.ColorDefault)

		PrintObject(shooter.Person)

		for _, bullet := range shooter.Bullets {
			PrintObject(*bullet)
		}

		for _, enemy := range enemies {
			PrintObject(enemy.Person)
		}

		term.Flush()
	}

}

func listenToKeyboard() {
	for {
		var event = term.PollEvent()

		if event.Type == term.EventKey {
			switch event.Key {
			case term.KeyArrowLeft:
				shooter.Person.MoveLeft()
			case term.KeyArrowRight:
				shooter.Person.MoveRight()
			case term.KeyArrowUp:
				shooter.Person.MoveUp()
			case term.KeyArrowDown:
				shooter.Person.MoveDown()
			case term.KeySpace:
				go shooter.Shoot(150)
			case term.KeyCtrlC:
				exit = true
			}
		}
	}
}

func GenerateEnemies() {
	ticker := time.NewTicker(time.Second * 5)

	for range ticker.C {
		width, height := term.Size()

		x := 0
		if rand.Float32() < 0.5 {
			x = width
		}

		y := 0
		if rand.Float32() < 0.5 {
			y = height
		}

		enemy := entities.Enemy{
			Person: entities.Object{Shape: '#', X: x, Y: y, Color: term.ColorRed},
			Target: &shooter.Person,
		}

		go enemy.GoKill(randomNumberBetween(8, 12))

		enemies = append(enemies, &enemy)
	}
}

func randomNumberBetween(min int, max int) int {
	return rand.Intn(max-min) + min
}

func PrintObject(object entities.Object) {
	term.SetCell(object.X, object.Y, object.Shape, object.Color, term.ColorDefault)
}
