package main

import (
	"math/rand"
)

type Board = [9][9]int

const EMPTY = 0

// NewBoard Generates a 9x9 Sudoku board.
func NewBoard() *Board {
	return generateBoard(nil, randomIntArray(), 0, 0)
}

// generateBoard Fills a 9x9 Sudoku board, starting at the specified row and column index.
// Accepts nil board to create an empty board.
func generateBoard(board *Board, sequence [9]int, row int, column int) *Board {
	// Create an empty board
	if board == nil {
		board = new(Board)
	}

	for _, num := range sequence {
		// Check if number can be played
		if !canPlay(board, num, row, column) {
			continue
		}

		// Play the number
		board[row][column] = num

		// Check what the following position to fill is (next column or next row)
		if column < 8 {
			generateBoard(board, sequence, row, column+1)
		} else if row < 8 {
			generateBoard(board, randomIntArray(), row+1, 0)
		}

		// If the board was filled successfully, finish
		if board[8][8] != EMPTY {
			return board
		}

		// If backtracking remove the one that was previously played to allow a new one to be played
		board[row][column] = EMPTY
	}

	return board
}

// canPlay Checks if a number you want to play can be played in that position following the Sudoku rules.
func canPlay(board *Board, number int, row int, column int) bool {
	// Checks if the number was played in the same row or column
	for index := range 9 {
		if board[row][index] == number || board[index][column] == number {
			return false
		}
	}

	// Search which 3x3 grid the number is in
	var rowRange, columnRange [3]int
	for start := 0; start < 9; start += 3 {
		end := start + 2

		if start <= row && row <= end {
			rowRange = [3]int{start, start + 1, end}
		}

		if start <= column && column <= end {
			columnRange = [3]int{start, start + 1, end}
		}
	}

	// Search 3x3 grid for number
	for _, thisRow := range rowRange {
		for _, thisColumn := range columnRange {
			if board[thisRow][thisColumn] == number {
				return false
			}
		}
	}

	return true
}

// NewBoardRemoveNumbers Remove n numbers from the board at random positions.
func NewBoardRemoveNumbers(board Board, n int) *Board {
	for i := 0; i < n; i++ {
		var x, y int
		// Gets a random cell position from the board until it finds a position that is not empty.
		for x, y = rand.Intn(9), rand.Intn(9); board[x][y] == EMPTY; {
			x, y = rand.Intn(9), rand.Intn(9)
		}
		board[x][y] = EMPTY
	}

	return &board
}

// randomIntArray Generates an array of 9 random integers from 1 to 9 (inclusive), and sorts it
// randomly.
func randomIntArray() [9]int {
	var trySequence [9]int
	for i := 1; i < 10; i++ {
		trySequence[i-1] = i
	}

	rand.Shuffle(9, func(i, j int) { trySequence[i], trySequence[j] = trySequence[j], trySequence[i] })
	return trySequence
}
