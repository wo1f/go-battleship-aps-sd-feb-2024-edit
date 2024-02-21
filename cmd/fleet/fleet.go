package fleet

import (
	"go-battleship/cmd/controller"
	"go-battleship/cmd/letter"
	"math/rand"
)

var fleets [][]*controller.Ship

func initializeEnemyFleet1() {
	fleet := controller.InitializeShips()

	fleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 4})
	fleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 5})
	fleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 6})
	fleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 7})
	fleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 8})

	fleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 6})
	fleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 7})
	fleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 9})

	fleet[2].SetPositions(&controller.Position{Column: letter.A, Row: 3})
	fleet[2].SetPositions(&controller.Position{Column: letter.B, Row: 3})
	fleet[2].SetPositions(&controller.Position{Column: letter.C, Row: 3})

	fleet[3].SetPositions(&controller.Position{Column: letter.F, Row: 8})
	fleet[3].SetPositions(&controller.Position{Column: letter.G, Row: 8})
	fleet[3].SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet[4].SetPositions(&controller.Position{Column: letter.C, Row: 5})
	fleet[4].SetPositions(&controller.Position{Column: letter.C, Row: 6})
}

func GetRandomFleet() []*controller.Ship {
	index := rand.Intn(len(fleets))
	return fleets[index]
}
