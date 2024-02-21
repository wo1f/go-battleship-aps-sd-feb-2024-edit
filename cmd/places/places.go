package places

import (
	"fmt"
	"go-battleship/cmd/controller"
	"go-battleship/cmd/letter"
	"math/rand"
)

var places [][]*controller.Ship

func place1() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 5})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 6})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 7})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 1})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 2})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 3})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 4})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.B, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.C, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 8})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 7})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.F, Row: 7})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 3})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.C, Row: 3})
	return fleet
}

func place2() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.B, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.D, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 3})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.B, Row: 5})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.B, Row: 6})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.B, Row: 7})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.B, Row: 8})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 5})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.E, Row: 5})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.F, Row: 5})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 1})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 2})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 3})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.G, Row: 7})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.G, Row: 8})
	return fleet
}

func place3() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 5})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 6})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 7})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.D, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.F, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.G, Row: 8})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 6})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 1})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 2})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 3})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.D, Row: 2})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.E, Row: 2})
	return fleet
}

func place4() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.A, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.A, Row: 2})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.A, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.A, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.A, Row: 5})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.F, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.G, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.C, Row: 5})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 5})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.E, Row: 5})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 2})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 3})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 4})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.C, Row: 2})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.D, Row: 2})
	return fleet
}

func place5() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.D, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.H, Row: 1})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.A, Row: 3})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.A, Row: 4})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.A, Row: 5})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.A, Row: 6})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 5})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 6})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 7})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.F, Row: 3})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 3})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 3})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.G, Row: 6})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.G, Row: 7})
	return fleet
}

func place6() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 2})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.G, Row: 5})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 5})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 6})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.A, Row: 1})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.B, Row: 1})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.C, Row: 1})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.C, Row: 3})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.D, Row: 3})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.E, Row: 3})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 6})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.C, Row: 6})
	return fleet
}

func place7() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.B, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.C, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.D, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.E, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 3})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.B, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.C, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.D, Row: 8})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 8})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 6})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.F, Row: 1})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.G, Row: 1})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 1})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 7})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.C, Row: 7})
	return fleet
}

func place8() *controller.Fleet {
	fleet := controller.NewFleet()

	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 1})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 2})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 3})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 4})
	fleet.AircraftCarrier.SetPositions(&controller.Position{Column: letter.F, Row: 5})

	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 1})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 2})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 3})
	fleet.Battleship.SetPositions(&controller.Position{Column: letter.E, Row: 4})

	fleet.Submarine.SetPositions(&controller.Position{Column: letter.D, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.E, Row: 8})
	fleet.Submarine.SetPositions(&controller.Position{Column: letter.F, Row: 8})

	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 6})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 7})
	fleet.Destroyer.SetPositions(&controller.Position{Column: letter.H, Row: 8})

	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 2})
	fleet.PatrolBoat.SetPositions(&controller.Position{Column: letter.B, Row: 3})
	return fleet
}

func fleetToShips(fleet *controller.Fleet) []*controller.Ship {
	ships := make([]*controller.Ship, 5)
	ships[0] = fleet.AircraftCarrier
	ships[1] = fleet.Battleship
	ships[2] = fleet.Submarine
	ships[3] = fleet.Destroyer
	ships[4] = fleet.PatrolBoat
	return ships
}

func init() {
	places = append(places, fleetToShips(place1()))
	places = append(places, fleetToShips(place2()))
	places = append(places, fleetToShips(place3()))
	places = append(places, fleetToShips(place4()))
	places = append(places, fleetToShips(place5()))
	places = append(places, fleetToShips(place6()))
	places = append(places, fleetToShips(place7()))
	places = append(places, fleetToShips(place8()))
}

func GetRandomPlace() []*controller.Ship {
	index := rand.Intn(len(places))
	fmt.Printf("place #%d (%d)", index, len(places))
	return places[index]
}
