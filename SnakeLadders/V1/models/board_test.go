package models

import (
	"testing"
)

func TestInitializeBoard(t *testing.T) {
	board := InitializeBoard()

	cell := board.GetCell(9, 0)
	if cell.GetCellNumber() != 1 {
		t.Errorf("Expected cell number %d, got %d at position (%d, %d)", 1, cell.GetCellNumber(), 9, 0)
	}

	cell = board.GetCell(9, 9)
	if cell.GetCellNumber() != 10 {
		t.Errorf("Expected cell number %d, got %d at position (%d, %d)", 10, cell.GetCellNumber(), 9, 9)
	}

	cell = board.GetCell(8, 9)
	if cell.GetCellNumber() != 11 {
		t.Errorf("Expected cell number %d, got %d at position (%d, %d)", 11, cell.GetCellNumber(), 8, 9)
	}

	cell = board.GetCell(0, 0)
	if cell.GetCellNumber() != 100 {
		t.Errorf("Expected cell number %d, got %d at position (%d, %d)", 100, cell.GetCellNumber(), 0, 0)
	}
}

func TestGetCell(t *testing.T) {
	board := InitializeBoard()

	// Test getting cells from different positions
	testCases := []struct {
		row, col int
		expected *Cell
	}{
		{0, 0, board.GetCell(0, 0)},
		{5, 5, board.GetCell(5, 5)},
		{9, 9, board.GetCell(9, 9)},
	}

	for _, tc := range testCases {
		if tc.expected != board.GetCell(tc.row, tc.col) {
			t.Errorf("Expected cell at position (%d, %d) to be %v, got %v", tc.row, tc.col, tc.expected, board.GetCell(tc.row, tc.col))
		}
	}
}

func TestSetCell(t *testing.T) {
	board := InitializeBoard()

	// Test setting cells to different values
	newCell := NewCell(999)
	board.SetCell(newCell, 0, 0)

	if board.GetCell(0, 0) != newCell {
		t.Errorf("Expected cell at position (0, 0) to be %v, got %v", newCell, board.GetCell(0, 0))
	}
}

func TestPrint(t *testing.T) {
	// For this test, we can check if the Print function runs without errors.
	// Manual verification is needed to ensure the output is correct.
	board := InitializeBoard()
	board.Print()
}
