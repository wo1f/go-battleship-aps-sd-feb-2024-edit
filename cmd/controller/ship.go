package controller

import (
	"fmt"
	"go-battleship/cmd/console"
	"go-battleship/cmd/letter"
	"strconv"
	"strings"
)

type Position struct {
	Column letter.Letter
	Row    int
}

func (p Position) String() string {
	return fmt.Sprintf("%s%d", p.Column.String(), p.Row)
}

type Fleet struct {
	AircraftCarrier *Ship
	Battleship      *Ship
	Submarine       *Ship
	Destroyer       *Ship
	PatrolBoat      *Ship
}

type Ship struct {
	isPlaced  bool
	Name      string
	Size      int
	positions []Position
	Color     console.Color
}

func NewPosition(letter letter.Letter, number int) *Position {
	return &Position{Column: letter, Row: number}
}

func (s *Ship) AddPosition(input string) {
	if s.positions == nil {
		s.positions = []Position{}
	}

	letter := letter.FromString(string(strings.ToUpper(input)[0]))
	number, err := strconv.Atoi(string(input[1]))

	if err != nil {
		panic(err)
	}

	s.positions = append(s.positions, Position{Column: letter, Row: number})
}

func (s *Ship) GetPositions() []Position {
	return s.positions
}

func (s *Ship) SetPositions(p *Position) {
	s.positions = append(s.positions, *p)
}
