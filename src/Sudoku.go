package main

import (
	"math/rand"
)

// Cell is an uint8 number that represents a number in a Sudoku cell.
type Cell uint8

// Row is a 9 long array of [Cell](s) that represent a single row in a Sudoku [Board].
type Row [9]Cell

// Board is a 9 long array of [Row](s) that represent a 9x9 Sudoku board.
type Board [9]Row

// EMPTY represents an empty sudoku cell.
const EMPTY Cell = 0

// difficulty represents how difficult a Sudoku board is.
//
// They signify the amount of empty cells that a board of a certain difficulty should have.
type difficulty uint8

// The difficulty settings supported.
//
// They start at 10 (easy) and increment by 10.
const (
	EASY difficulty = 10 + (iota * 10)
	MEDIUM
	HARD
)

// NewBoard generates a random [Board].
func NewBoard() *Board {
	return generateBoard(nil, shuffledRow(), 0, 0)
}

// generateBoard fills a [Board], starting at the specified row and column index.
// Accepts a nil board, which creates an empty [Board].
func generateBoard(board *Board, sequence *Row, row int, column int) *Board {
	// Create an empty board
	if board == nil {
		board = new(Board)
	}

	for _, num := range sequence {
		// Check if number can be played
		if !isValidCell(board, num, row, column) {
			continue
		}

		// Play the number
		board[row][column] = num

		// Check what the following position to fill is (next column or next row)
		if column < 8 {
			generateBoard(board, sequence, row, column+1)
		} else if row < 8 {
			generateBoard(board, shuffledRow(), row+1, 0)
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

// isValidCell reports whether a number can be played (according to the Sudoku rules) in a specific
// row and column position.
func isValidCell(board *Board, number Cell, row int, column int) bool {
	// Disallow inserting in an already occupied cell
	if board[row][column] != EMPTY {
		return false
	}

	// Checks if the number was played in the same row or column
	for index := range 9 {
		if board[row][index] == number || board[index][column] == number {
			return false
		}
	}

	// Search which 3x3 grid that the number is in
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

	// Search the 3x3 grid for the number
	for _, thisRow := range rowRange {
		for _, thisColumn := range columnRange {
			if board[thisRow][thisColumn] == number {
				return false
			}
		}
	}

	return true
}

// NewBoardForDifficulty generates a random unsolved [Board] based on the difficulty desired.
func NewBoardForDifficulty(diff difficulty) *Board {
	var board = NewBoard()

	for i := 0; i < int(diff); i++ {
		var x, y int
		// Gets a random cell position from the board until it finds a position that is not empty.
		for x, y = rand.Intn(9), rand.Intn(9); board[x][y] == EMPTY; {
			x, y = rand.Intn(9), rand.Intn(9)
		}
		board[x][y] = EMPTY
	}

	return board
}

// shuffledRow generates a valid Sudoku [Row] with a random/shuffled order.
func shuffledRow() *Row {
	var row = new(Row)
	for i := 0; i < 9; i++ {
		row[i] = Cell(i + 1)
	}

	rand.Shuffle(9, func(i, j int) {
		row[i], row[j] = row[j], row[i]
	})

	return row
}
