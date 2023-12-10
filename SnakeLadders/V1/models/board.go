package models

import "fmt"

type Board struct {
	board   [10][10]*Cell
	snakes  map[int]int // map of start to end
	ladders map[int]int // map of start to end
}

func InitializeBoard() *Board {
	board := [10][10]*Cell{}

	currVal := 1
	leftToright := 1
	for r := 9; r >= 0; r-- {
		for c := 0; c < 10; c++ {
			if leftToright == 1 {
				board[r][c] = NewCell(currVal)
			} else {
				board[r][9-c] = NewCell(currVal)
			}
			currVal++
		}
		leftToright = 1 - leftToright
	}

	snakes := make(map[int]int)
	ladders := make(map[int]int)

	return &Board{
		board:   board,
		snakes:  snakes,
		ladders: ladders,
	}
}

func (b *Board) GetCell(r, c int) *Cell {
	return b.board[r][c]
}

func (b *Board) SetCell(cell *Cell, r, c int) *Board {
	b.board[r][c] = cell
	return b
}

func (b *Board) Print() {

	for i := 0; i < 10; i++ {
		fmt.Println("============================================================")
		for j := 0; j < 10; j++ {
			fmt.Printf("%03d | ", b.GetCell(i, j).GetCellNumber())
		}
		fmt.Println()
	}
}

func (b *Board) AddSnake(start, end *Cell) {
	b.snakes[start.GetCellNumber()] = end.GetCellNumber()
}

func (b *Board) AddLadder(start, end *Cell) {
	b.snakes[start.GetCellNumber()] = end.GetCellNumber()
}

func (b *Board) IsSnake(cell *Cell) bool {
	if _, ok := b.snakes[cell.GetCellNumber()]; !ok {
		return false
	}
	return true
}

func (b *Board) GetSnakeDest(cell *Cell) *Cell {
	return NewCell(b.snakes[cell.GetCellNumber()])
}

func (b *Board) IsLadder(cell *Cell) bool {
	if _, ok := b.ladders[cell.GetCellNumber()]; !ok {
		return false
	}
	return true
}

func (b *Board) GetLadderDest(cell *Cell) *Cell {
	return NewCell(b.ladders[cell.GetCellNumber()])
}
