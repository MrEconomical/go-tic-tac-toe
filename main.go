package main

// Imports

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	turn := uint(1)

	// Run game loop

	reader := bufio.NewReader(os.Stdin)
	clearScreen()
	fmt.Println("Welcome to Bad Tic Tac Toe!")

	for {
		// Get symbol and display instructions

		var symbol rune
		if turn == 1 {
			symbol = 'X'
		} else {
			symbol = 'O'
		}

		displayBoard(&board)
		fmt.Printf("Player %[1]v's turn (place an %[1]v)\n", string(symbol))
		fmt.Println("Enter a square to make a move:")

		// Parse player input

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		row, col, valid := parseInput(input)

		if !valid || board[row][col] != 0 {
			clearScreen()
			fmt.Printf("\"%v\" is not a valid input square, please try again (i.e. a1)\n", input)
			continue
		}

		// Make move and check result

		board[row][col] = turn
		result, end := checkResult(&board)
		if end {
			clearScreen()
			if result == 0 {
				fmt.Println("Game drawn!")
			} else if result == 1 {
				fmt.Println("Player X wins!")
			} else {
				fmt.Println("Player O wins!")
			}
			displayBoard(&board)
			break
		}

		clearScreen()
		if turn == 1 {
			turn = 2
		} else {
			turn = 1
		}
	}
}

// Get row and column from player input

func parseInput(input string) (uint, uint, bool) {
	if len(input) < 2 ||
		(input[0] != 'a' && input[0] != 'b' && input[0] != 'c') ||
		(input[1] != '1' && input[1] != '2' && input[1] != '3') {
		return 0, 0, false
	}
	return uint(input[0]) - 97, uint(input[1]) - 49, true
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
		fmt.Print(rowLabel(r) + " ")
		for _, s := range row {
			if s == 0 {
				fmt.Print("_ ")
			} else if s == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print("O ")
			}
		}
		fmt.Println(rowLabel(r))
	}
	fmt.Println("  1 2 3")
}

// Get row label from index

func rowLabel(index int) string {
	if index == 0 {
		return "a"
	} else if index == 1 {
		return "b"
	}
	return "c"
}

// Clear terminal screen

func clearScreen() {
	fmt.Println("\033[2J\033[H")
}
