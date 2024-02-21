package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"go-battleship/cmd/console"
	"go-battleship/cmd/controller"
	"go-battleship/cmd/letter"
)

var (
	enemyFleet []*controller.Ship
	myFleet    []*controller.Ship
	scanner    *bufio.Scanner
)
var printer = console.ColoredPrinter(1, false).Background(console.BLACK).Foreground(console.WHITE).Build()

func run() {
	scanner = bufio.NewScanner(os.Stdin)

	printer.SetForegroundColor(console.MAGENTA)
	printer.Println("                                     |__")
	printer.Println("                                     |\\/")
	printer.Println("                                     ---")
	printer.Println("                                     / | [")
	printer.Println("                              !      | |||")
	printer.Println("                            _/|     _/|-++'")
	printer.Println("                        +  +--|    |--|--|_ |-")
	printer.Println("                     { /|__|  |/\\__|  |--- |||__/")
	printer.Println("                    +---------------___[}-_===_.'____                 /\\")
	printer.Println("                ____`-' ||___-{]_| _[}-  |     |_[___\\==--            \\/   _")
	printer.Println(" __..._____--==/___]_|__|_____________________________[___\\==--____,------' .7")
	printer.Println("|                        Welcome to Battleship                         BB-61/")
	printer.Println(" \\_________________________________________________________________________|")
	printer.Println("")
	printer.SetForegroundColor(console.WHITE)

	initializeGame()

	startGame()
}

func main() {
	run()
}

func myShot() {
	printer.Println("")
	printer.Println("Player, it's your turn")
	printer.Println("Enter coordinates for your shot (i.e A3) :")
	var isHit bool
	for scanner.Scan() {
		input := scanner.Text()
		if input != "" {
			position := parsePosition(input)
			var err error
			isHit, err = controller.CheckIsHit(enemyFleet, position)
			if err != nil {
				printer.Printf("Error: %s\n", err)
			}
			fmt.Println(enemyFleet[4].IsAlive())
			break
		}
	}

	resultOfShot(isHit)

}

func enemyShot() {
	var isHit bool
	position := getRandomPosition()
	var err error
	isHit, err = controller.CheckIsHit(myFleet, position)
	if err != nil {
		printer.Printf("Error: %s\n", err)
	}
	printer.Println("")

	result := "hit your ship !"
	printer.SetForegroundColor(HitColor)
	if !isHit {
		printer.SetForegroundColor(MissColor)
		result = "miss"
	}

	printer.Printf("Computer shoot in %s%d and %s\n", position.Column, position.Row, result)
	printer.SetForegroundColor(planeColor)
	if isHit {
		drawExplosion()
	}

}

func drawState() {
	var destroyedShips []*controller.Ship
	var aliveShips []*controller.Ship
	for _, ship := range enemyFleet {
		if !ship.IsAlive() {
			destroyedShips = append(destroyedShips, ship)
		} else {
			aliveShips = append(aliveShips, ship)
		}
	}
	if len(destroyedShips) != 0 {
		printer.Printf("You have destroyed the following ships:\n")
		for _, ship := range destroyedShips {
			printer.SetForegroundColor(ship.Color)
			printer.Printf("%s (%d holes)\n", ship.Name, ship.Size)
		}
		printer.SetForegroundColor(planeColor)
		printer.Println("==============================================")
	}
	if len(aliveShips) != 0 {
		printer.Printf("The enemy still has ships:\n")
		for _, ship := range aliveShips {
			printer.SetForegroundColor(ship.Color)
			printer.Printf("%s (%d holes)\n", ship.Name, ship.Size)

		}
		printer.SetForegroundColor(planeColor)
	}
}

func isEndGame(ships []*controller.Ship) bool {
	for _, ship := range ships {
		if ship.IsAlive() {
			return false
		}
	}
	return true
}

func round(number int) {
	printer.Printf("==============================================\n")
	printer.Printf("Round %d\n", number)
	printer.Printf("==============================================\n")

	drawState()

	myShot()
	if isEndGame(enemyFleet) {
		printer.SetForegroundColor(console.RED)
		printer.Println("You Win!")
		os.Exit(0)
	}

	enemyShot()
	if isEndGame(myFleet) {
		printer.SetForegroundColor(console.RED)
		printer.Println("You Lose!")
		os.Exit(0)
	}
}

func startGame() {
	printer.Print("\033[2J\033[;H")
	printer.Println("                  __")
	printer.Println("                 /  \\")
	printer.Println("           .-.  |    |")
	printer.Println("   *    _.-'  \\  \\__/")
	printer.Println("    \\.-'       \\")
	printer.Println("   /          _/")
	printer.Println("  |      _  /\" \"")
	printer.Println("  |     /_\\'")
	printer.Println("   \\    \\_/")
	printer.Println("    \" \"\" \"\" \"\" \"")

	i := 1
	for {
		round(i)
		i++
	}
}

func parsePosition(input string) *controller.Position {
	ltr := strings.ToUpper(string(input[0]))
	number, _ := strconv.Atoi(string(input[1]))
	return &controller.Position{
		Column: letter.FromString(ltr),
		Row:    number,
	}
}

func beep() {
	fmt.Print("\007")
}

func initializeGame() {
	initializeMyFleet()

	//initializeMyFleetConst()

	initializeEnemyFleet()
}

func initializeMyFleetConst() {
	myFleet = controller.InitializeShips()

	myFleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 4})
	myFleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 5})
	myFleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 6})
	myFleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 7})
	myFleet[0].SetPositions(&controller.Position{Column: letter.B, Row: 8})

	myFleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 6})
	myFleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 7})
	myFleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 8})
	myFleet[1].SetPositions(&controller.Position{Column: letter.E, Row: 9})

	myFleet[2].SetPositions(&controller.Position{Column: letter.A, Row: 3})
	myFleet[2].SetPositions(&controller.Position{Column: letter.B, Row: 3})
	myFleet[2].SetPositions(&controller.Position{Column: letter.C, Row: 3})

	myFleet[3].SetPositions(&controller.Position{Column: letter.F, Row: 8})
	myFleet[3].SetPositions(&controller.Position{Column: letter.G, Row: 8})
	myFleet[3].SetPositions(&controller.Position{Column: letter.H, Row: 8})

	myFleet[4].SetPositions(&controller.Position{Column: letter.C, Row: 5})
	myFleet[4].SetPositions(&controller.Position{Column: letter.C, Row: 6})
}

func initializeMyFleet() {
	//reader := bufio.NewReader(os.Stdin)
	//scanner := bufio.NewScanner(os.Stdin)
	myFleet = controller.InitializeShips()

	printer.Println("Please position your fleet (Game board has size from A to H and 1 to 8) :")

	for _, ship := range myFleet {
		printer.Println("")
		printer.Printf("Please enter the positions for the %s (size: %d)", ship.Name, ship.Size)
		printer.Println("")

		for i := 1; i <= ship.Size; i++ {
			printer.Printf("Enter position %d of %d (i.e A3):\n", i, ship.Size)

			for scanner.Scan() {
				positionInput := scanner.Text()
				if positionInput != "" {
					ship.AddPosition(positionInput)
					break
				}
			}
		}
	}
}

func getRandomPosition() *controller.Position {
	rows := 8
	lines := 8
	letter := letter.Letter(rand.Intn(lines-1) + 1)
	number := rand.Intn(rows-1) + 1
	return &controller.Position{Column: letter, Row: number}
}

const HitColor = console.ORANGE
const MissColor = console.CADET_BLUE
const planeColor = console.WHITE
const explosionColor = console.RED

func drawExplosion() {
	beep()
	printer.SetForegroundColor(explosionColor)
	printer.Println("                \\         .  ./")
	printer.Println("              \\      .:\" \";'.:..\" \"   /")
	printer.Println("                  (M^^.^~~:.'\" \").")
	printer.Println("            -   (/  .    . . \\ \\)  -")
	printer.Println("               ((| :. ~ ^  :. .|))")
	printer.Println("            -   (\\- |  \\ /  |  /)  -")
	printer.Println("                 -\\  \\     /  /-")
	printer.Println("                   \\  \\   /  /")
	printer.SetForegroundColor(planeColor)
}

func resultOfShot(isHit bool) {
	if isHit {
		drawExplosion()
		printer.SetForegroundColor(HitColor)
		printer.Println("Yeah ! Nice hit !")
	} else {
		printer.SetForegroundColor(MissColor)
		printer.Println("Miss")
	}
	printer.SetForegroundColor(planeColor)
}
