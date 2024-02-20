package controller

import (
	"github.com/stretchr/testify/assert"
	"go-battleship/cmd/console"
	"testing"
)

func IsAliveTest(t testing.T) {
	t.Run("new ship", func(t *testing.T) {
		ship := NewShip("Patrol Boat", 2, console.ORANGE)

		assert.True(t, ship.IsAlive())
	})
	t.Run("new dead ship", func(t *testing.T) {
		ship := NewShip("Patrol Boat", 2, console.ORANGE)
		for _, pos := range ship.positions {
			pos.Hitted = true
		}

		assert.False(t, ship.IsAlive())
	})

}
