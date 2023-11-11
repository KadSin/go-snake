package tests

import (
	"kadsin/shoot-run/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var obj entities.Object

func reset() {
	entities.TerminalSize = func() (int, int) { return 20, 20 }

	obj.X = 10
	obj.Y = 10
}

func TestUpdateLocation(t *testing.T) {
	cases := map[string][2]func(){
		"Up":    {obj.MoveUp, func() { assert.Equal(t, 5, obj.Y) }},
		"Right": {obj.MoveRight, func() { assert.Equal(t, 15, obj.X) }},
		"Down":  {obj.MoveDown, func() { assert.Equal(t, 15, obj.Y) }},
		"Left":  {obj.MoveLeft, func() { assert.Equal(t, 5, obj.X) }},
	}

	for name, assertion := range cases {
		t.Run(name, func(t *testing.T) {
			reset()

			assertion[0]()
			obj.UpdateLocation(5)

			assertion[1]()
		})
	}
}

func TestGetErrorWhenExceedsBoundaries(t *testing.T) {
	cases := map[string]func(){
		"Top":    obj.MoveUp,
		"Right":  obj.MoveRight,
		"Bottom": obj.MoveDown,
		"Left":   obj.MoveLeft,
	}

	for name, changeDirection := range cases {
		t.Run(name, func(t *testing.T) {
			reset()

			changeDirection()
			error := obj.UpdateLocation(11)

			assert.Error(t, error)
		})
	}
}
