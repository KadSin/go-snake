package tests

import (
	"kadsin/shoot-run/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var obj entities.Object

func TestUpdateLocationToUp(t *testing.T) {
	obj.MoveUp()
	UpdateLocation()

	assert.Equal(t, 5, obj.Y)
}

func TestUpdateLocationToRight(t *testing.T) {
	obj.MoveRight()
	UpdateLocation()

	assert.Equal(t, 15, obj.X)
}

func TestUpdateLocationToDown(t *testing.T) {
	obj.MoveDown()
	UpdateLocation()

	assert.Equal(t, 15, obj.Y)
}

func TestUpdateLocationToLeft(t *testing.T) {
	obj.MoveLeft()
	UpdateLocation()

	assert.Equal(t, 5, obj.X)
}

func UpdateLocation() {
	entities.TerminalSize = func() (int, int) { return 20, 20 }

	obj.X = 10
	obj.Y = 10

	obj.UpdateLocation(5)
}
