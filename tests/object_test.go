package tests

import (
	"kadsin/shoot-run/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var obj = entities.Object{
	ScreenStart: entities.Coordinate{X: 0, Y: 0},
	ScreenSize:  entities.Coordinate{X: 20, Y: 20},
}

func reset() {
	obj.Location = entities.Coordinate{X: 10, Y: 10}
}

func TestUpdateLocation(t *testing.T) {
	t.Run("Up", func(t *testing.T) {
		reset()
		obj.MoveUp()
		obj.UpdateLocation(5)

		assert.Equal(t, 5, obj.Location.Y)
	})

	t.Run("Right", func(t *testing.T) {
		reset()
		obj.MoveRight()
		obj.UpdateLocation(5)

		assert.Equal(t, 15, obj.Location.X)
	})

	t.Run("Down", func(t *testing.T) {
		reset()
		obj.MoveDown()
		obj.UpdateLocation(5)

		assert.Equal(t, 15, obj.Location.Y)
	})

	t.Run("Left", func(t *testing.T) {
		reset()
		obj.MoveLeft()
		obj.UpdateLocation(5)

		assert.Equal(t, 5, obj.Location.X)
	})
}

func TestGetErrorWhenExceedsBoundaries(t *testing.T) {
	reset()

	t.Run("Top", func(t *testing.T) {
		obj.MoveUp()
		error := obj.UpdateLocation(11)

		assert.Error(t, error)
	})

	t.Run("Right", func(t *testing.T) {
		reset()
		obj.MoveRight()
		error := obj.UpdateLocation(11)

		assert.Error(t, error)
	})

	t.Run("Bottom", func(t *testing.T) {
		reset()
		obj.MoveDown()
		error := obj.UpdateLocation(11)

		assert.Error(t, error)
	})

	t.Run("Left", func(t *testing.T) {
		reset()
		obj.MoveLeft()
		error := obj.UpdateLocation(11)

		assert.Error(t, error)
	})
}
