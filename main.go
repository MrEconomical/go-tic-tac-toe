package main

// Imports

import "fmt"

// Run Tic Tac Toe game

func main() {
	// Initialize board and turn counter
	// Board square: 0 = empty, 1 = player X, 2 = player O
	// Turn: 1 = player X, 2 = player O

	board := [3][3]uint{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	//turn := 1

	displayBoard(&board)
}

// Check for player win or draw
// 0 = no result, 1/2 = win, 3 = draw

func checkResult(board *[3][3]uint) uint {
	// Check rows

	for _, row := range board {
		if r := checkValues(row[0], row[1], row[2]); r != 0 {
			return r
		}
	}

	return 0
}

// Check for win from three values
// 0 = no result, 1/2 = win

func checkValues(a, b, c uint) uint {
	if a == b && b == c {
		return a
	}
	return 0
}

// Print board to terminal with row and column labels

func displayBoard(board *[3][3]uint) {
	fmt.Println("  1 2 3")
	for r, row := range board {
		fmt.Printf("%v ", r+1)
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
