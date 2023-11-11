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
	t.Run("Up", func(t *testing.T) {
		reset()
		obj.MoveUp()
		obj.UpdateLocation(5)

		assert.Equal(t, 5, obj.Y)
	})

	t.Run("Right", func(t *testing.T) {
		reset()
		obj.MoveRight()
		obj.UpdateLocation(5)

		assert.Equal(t, 15, obj.X)
	})

	t.Run("Down", func(t *testing.T) {
		reset()
		obj.MoveDown()
		obj.UpdateLocation(5)

		assert.Equal(t, 15, obj.Y)
	})

	t.Run("Left", func(t *testing.T) {
		reset()
		obj.MoveLeft()
		obj.UpdateLocation(5)

		assert.Equal(t, 5, obj.X)
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
