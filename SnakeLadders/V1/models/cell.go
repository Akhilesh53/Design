package models

import "fmt"

type Cell struct {
	number int
}

func NewCell(number int) *Cell {
	return &Cell{
		number: number,
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
