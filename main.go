package main

// Imports

import (
	"bufio"
	"fmt"
	"os"
)

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
	turn := 1

	clearScreen()
	fmt.Println("Welcome to Bad Tic Tac Toe!")
	displayBoard(&board)

	// Run game loop

	reader := bufio.NewReader(os.Stdin)

	for {
		// Get symbol and display instructions

		var symbol rune
		if turn == 1 {
			symbol = 'X'
		} else {
			symbol = 'O'
		}

		fmt.Printf("Player %[1]v's turn (place an %[1]v)\n", string(symbol))
		fmt.Println("Enter a square to make a move:")

		// Parse player input

		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}
}

// Check for player win or draw
// 0 = draw, 1/2 = win

func checkResult(board *[3][3]uint) (uint, bool) {
	// Check rows and columns

	for r := 0; r < 3; r++ {
		if result := checkValues(board[r][0], board[r][1], board[r][2]); result != 0 {
			return result, true
		}
	}

	for c := 0; c < 3; c++ {
		if result := checkValues(board[0][c], board[1][c], board[2][c]); result != 0 {
			return result, true
		}
	}

	// Check diagonals

	if result := checkValues(board[0][0], board[1][1], board[2][2]); result != 0 {
		return result, true
	}
	if result := checkValues(board[2][0], board[1][1], board[0][2]); result != 0 {
		return result, true
	}

	// Check draw if board full otherwise no result

	for _, row := range board {
		for _, s := range row {
			if s == 0 {
				return 0, false
			}
		}
	}

	return 0, true
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
