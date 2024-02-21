package places

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-battleship/cmd/controller"
	"go-battleship/cmd/letter"
	"testing"
)

func placeBad() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: -1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 2})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 10})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 5})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 1})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 0})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 3})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 4})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.F, Row: 8})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 6})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 2})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 3})
	return fleet
}

func TestPlace(t *testing.T) {
	for i := range places {
		t.Run(fmt.Sprintf("place %d", i), func(t *testing.T) {
			settedPos := make(map[string]struct{})
			for _, ship := range places[i] {
				for _, pos := range ship.GetPositions() {
					_, ok := settedPos[pos.String()]
					assert.Falsef(t, ok, "collision")
					settedPos[pos.String()] = struct{}{}

					assert.LessOrEqualf(t, pos.Row, 8, "row > 8")
					assert.GreaterOrEqualf(t, pos.Row, 1, "row < 1")
				}
			}
		})
	}
}
