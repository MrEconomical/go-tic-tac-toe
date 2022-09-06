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

	fmt.Println(board)
}

// Clear terminal screen

func clearScreen() {
	fmt.Println("\033[2J\033[H")
}