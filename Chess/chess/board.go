package chess

import "fmt"

type Board struct {
	size  int
	board [][]ISpot
}

func InitialiseBoard(sz int) *Board {
	board := make([][]ISpot, sz)
	for i := 0; i < sz; i++ {
		board[i] = make([]ISpot, sz)
	}
	return &Board{size: sz, board: board}
}

func (b *Board) AddSpot(spot ISpot) {
	b.board[spot.GetPosition().GetX()][spot.GetPosition().GetY()] = spot
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
				fmt.Print(b.GetSpot(i, j).GetPiece().GetPieceType().String())
			}

			fmt.Print(" | ")
		}
		fmt.Println()
	}
}

func (b *Board) GetSpot(row, col int) ISpot {
	return b.board[row][col]
}
