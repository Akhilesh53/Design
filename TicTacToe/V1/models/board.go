package models

import "fmt"

type Board struct {
	size  int
	board [][]PlayingPiece
}

func InitialiseBoard(sz int) *Board {
	board := make([][]PlayingPiece, sz)
	for i := 0; i < sz; i++ {
		board[i] = make([]PlayingPiece, sz)
	}
	return &Board{size: sz, board: board}
}

func (b *Board) AddPiece(piece PlayingPiece, row, col int) {
	b.board[row][col] = piece
}

func (b *Board) IsNull(row, col int) bool {
	return b.board[row][col] == nil
}

func (b *Board) PrintBoard() {
	fmt.Println()
	fmt.Println("=========== Board ===============")
	defer fmt.Println("=================================")
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.board[i][j] == nil {
				fmt.Print("   ")
			} else {
				fmt.Print(" " + (b.board[i][j]).GetPiece() + " ")
			}

			fmt.Print(" | ")
		}
		fmt.Println()
	}
}

func (b *Board) GetVacantPlaces() int {
	cnt := 0
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			if b.IsNull(i, j) {
				cnt++
			}
		}
	}
	return cnt
}

func (b *Board) GetPiece(row, col int) PlayingPiece {
	return b.board[row][col]
}

func (b *Board) IsPieceInAnyRow(piece PlayingPiece) bool {
	for i := 0; i < b.size; i++ {
		inRow := true
		for j := 0; j < b.size; j++ {
			if b.IsNull(i, j) || b.GetPiece(i, j) != piece {
				inRow = false
				break
			}
		}
		if !inRow {
			continue
		}
		return true
	}
	return false
}

func (b *Board) IsPieceInAnyColumn(piece PlayingPiece) bool {
	for i := 0; i < b.size; i++ {
		inCol := true
		for j := 0; j < b.size; j++ {
			if b.IsNull(j, i) || b.GetPiece(j, i) != piece {
				inCol = false
				break
			}
		}
		if !inCol {
			continue
		}
		return true
	}
	return false
}

func (b *Board) IsPieceInAnyDiagonal(piece PlayingPiece) bool {
	inDiagonal := true
	for i := 0; i < b.size; i++ {
		if b.IsNull(i, i) || b.GetPiece(i, i) != piece {
			inDiagonal = false
		}
	}

	if inDiagonal {
		return true
	}
	inDiagonal = true
	for i := 0; i < b.size; i++ {
		if b.IsNull(i, b.size-i-1) || b.GetPiece(i, b.size-i-1) != piece {
			inDiagonal = false
		}
	}
	return inDiagonal
}
