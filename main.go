package main

import "fmt"

func main() {
	// Initialize board and turn counter
	// Board square: 0 = Empty, 1 = Player X, 2 = Player O
	// Turn: 1 = Player X, 2 = Player O

	board := [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	//turn := 1

	displayBoard(&board)
}

// Print board to terminal with row and column labels

func displayBoard(board *[3][3]int) {
	fmt.Println("  1 2 3")
	for r, row := range board {
		fmt.Printf("%v ", r + 1)
		for _, s := range row {
			if s == 0 {
				fmt.Print("_ ")
			} else if s == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println(r + 1)
	}
	fmt.Println("  1 2 3")
}

// Clear terminal screen

func clearScreen() {
	fmt.Println("\033[2J\033[H")
}