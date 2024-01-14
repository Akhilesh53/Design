package models

import "fmt"

type Board struct {
	board   [10][10]*Cell
	snakes  *Snake
	ladders *Ladder
}

func InitializeBoard() *Board {
	board := [10][10]*Cell{}
	snakes := MapSnakes()
	ladders := MapLadders()

	leftToRight := 1
	for r := 9; r >= 0; r-- {
		for c := 0; c < 10; c++ {
			currVal := 10 * (9 - r)
			col := 0
			if leftToRight == 1 {
				col = c
			} else {
				col = 9 - c
			}
			currVal += col + 1
			fmt.Print(r, ",", c, "->", currVal, ", ")
			board[r][col] = NewCell(currVal).SetSnake(snakes.IsSnake(currVal)).SetLadder(ladders.IsLadder(currVal))
		}
		fmt.Println()
		leftToRight = 1 - leftToRight
	}

	return &Board{
		board:   board,
		snakes:  snakes,
		ladders: ladders,
	}
}

func (b *Board) GetCell(r, c int) *Cell {
	return b.board[r][c]
}

func (b *Board) Print() {
	leftToRight := 0
	for i := 0; i < 10; i++ {
		if leftToRight == 1 {
			for j := 0; j < 10; j++ {
				fmt.Printf("%03d | ", b.GetCell(i, j).GetCellNumber())
			}
		} else {
			for j := 9; j >= 0; j-- {
				fmt.Printf("%03d | ", b.GetCell(i, j).GetCellNumber())
			}
		}
		fmt.Println()
		leftToRight = 1 - leftToRight
	}
}

func (b *Board) AddSnake(start, end *Cell) {
	b.snakes.AddSnake(start.GetCellNumber(), end.GetCellNumber())
}

func (b *Board) AddLadder(start, end *Cell) {
	b.ladders.AddLadder(start.GetCellNumber(), end.GetCellNumber())
}

func (b *Board) GetSnakesMap() *Snake {
	return b.snakes
}

func (b *Board) GetLaddersMap() *Ladder {
	return b.ladders
}

func (b *Board) GetCellByCellNumber(cellNumber int) *Cell {
	row := 9 - ((cellNumber - 1) / 10)
	var column int

	if row%2 == 0 {
		column = 9 - ((cellNumber - 1) % 10)
	} else {
		column = (cellNumber - 1) % 10
	}
	fmt.Println(row, column, b.GetCell(row, column))
	return b.GetCell(row, column)
}
