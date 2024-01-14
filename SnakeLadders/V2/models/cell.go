package models

import "fmt"

type Cell struct {
	number   int
	isSnake  bool
	isLadder bool
}

func NewCell(number int) *Cell {
	return &Cell{
		number:   number,
		isSnake:  false,
		isLadder: false,
	}
}

func NewDefaultCell() *Cell {
	return &Cell{}
}

func (c *Cell) GetCellNumber() int {
	return c.number
}

func (c *Cell) String() string {
	return fmt.Sprintf("%d", c.number)
}

func (c *Cell) SetCellNumber(number int) *Cell {
	c.number = number
	return c
}

func (c *Cell) IsSnake() bool  { return c.isSnake }
func (c *Cell) IsLadder() bool { return c.isLadder }

func (c *Cell) SetSnake(val bool) *Cell {
	c.isSnake = val
	return c
}

func (c *Cell) SetLadder(val bool) *Cell {
	c.isLadder = val
	return c
}
