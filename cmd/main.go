package main

import (
	"bufio"
	"fmt"
	"go-battleship/cmd/places"
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

func main() {

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

	for {
		printer.Println("")
		printer.Println("Player, it's your turn")
		printer.Println("Enter coordinates for your shot :")
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
				break
			}
		}

		if isHit {
			beep()
			printer.Println("                \\         .  ./")
			printer.Println("              \\      .:\" \";'.:..\" \"   /")
			printer.Println("                  (M^^.^~~:.'\" \").")
			printer.Println("            -   (/  .    . . \\ \\)  -")
			printer.Println("               ((| :. ~ ^  :. .|))")
			printer.Println("            -   (\\- |  \\ /  |  /)  -")
			printer.Println("                 -\\  \\     /  /-")
			printer.Println("                   \\  \\   /  /")
		}

		if isHit {
			printer.Println("Yeah ! Nice hit !")
		} else {
			printer.Println("Miss")
		}

		position := getRandomPosition()
		var err error
		isHit, err = controller.CheckIsHit(myFleet, position)
		if err != nil {
			printer.Printf("Error: %s\n", err)
		}
		printer.Println("")

		result := "hit your ship !"
		if !isHit {
			result = "miss"
		}
		printer.Printf("Computer shoot in %s%d and %s\n", position.Column, position.Row, result)
		if isHit {
			beep()
			printer.Println("                \\         .  ./")
			printer.Println("              \\      .:\" \";'.:..\" \"   /")
			printer.Println("                  (M^^.^~~:.'\" \").")
			printer.Println("            -   (/  .    . . \\ \\)  -")
			printer.Println("               ((| :. ~ ^  :. .|))")
			printer.Println("            -   (\\- |  \\ /  |  /)  -")
			printer.Println("                 -\\  \\     /  /-")
			printer.Println("                   \\  \\   /  /")
		}
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

	//myFleet = places.GetRandomPlace()

	enemyFleet = places.GetRandomPlace()
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
