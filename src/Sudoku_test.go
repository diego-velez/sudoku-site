package main

import (
	"fmt"
	"testing"
)

func TestIsValidCell(t *testing.T) {
	type test struct {
		board    *Board
		row      int
		column   int
		number   Cell
		expected bool
	}

	var board = &Board{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	var tests = []test{
		// Test inserting in occupied cells
		{board, 0, 0, 1, false},
		{board, 8, 8, 1, false},
		{board, 5, 4, 3, false},
		{board, 4, 3, 5, false},

		// Test 3x3 grid
		{board, 4, 4, 8, false},
		{board, 7, 7, 5, false},
		{board, 8, 0, 6, false},
		{board, 1, 1, 9, false},

		// Test column
		{board, 3, 3, 1, false},
		{board, 6, 4, 2, false},
		{board, 8, 0, 5, false},
		{board, 6, 4, 7, false},

		// Test row
		{board, 3, 2, 3, false},
		{board, 8, 3, 9, false},
		{board, 0, 8, 7, false},
		{board, 7, 7, 4, false},

		// Test numbers outside the [1-9] range
		{board, 3, 2, 0, false},
		{board, 8, 3, 10, false},
		{board, 0, 8, 69, false},
		{board, 7, 7, 156, false},

		// Test valid
		{board, 0, 2, 1, true},
		{board, 0, 5, 2, true},
		{board, 5, 2, 3, true},
		{board, 0, 7, 9, true},
	}

	for i, test := range tests {
		name := fmt.Sprintf(
			"%d out of %d (%d,%d) with %d",
			i,
			len(tests)-1,
			test.row,
			test.column,
			test.number,
		)

		t.Run(name, func(t *testing.T) {
			if isValidCell(test.board, test.number, test.row, test.column) != test.expected {
				t.Errorf("Expected %t but got %t", test.expected, !test.expected)
			}
		})
	}

	t.Run("test non-destructive", func(t *testing.T) {
		var boardCopy = Board{
			{5, 3, 0, 0, 7, 0, 0, 0, 0},
			{6, 0, 0, 1, 9, 5, 0, 0, 0},
			{0, 9, 8, 0, 0, 0, 0, 6, 0},
			{8, 0, 0, 0, 6, 0, 0, 0, 3},
			{4, 0, 0, 8, 0, 3, 0, 0, 1},
			{7, 0, 0, 0, 2, 0, 0, 0, 6},
			{0, 6, 0, 0, 0, 0, 2, 8, 0},
			{0, 0, 0, 4, 1, 9, 0, 0, 5},
			{0, 0, 0, 0, 8, 0, 0, 7, 9},
		}

		if *board != boardCopy {
			t.Error("Function is modifying the board")
		}
	})
}
