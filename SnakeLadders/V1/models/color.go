package models

import "fmt"

type Color int

const (
	RED Color = iota
	BLUE
	YELLOW
	GREEN
)

func (c *Color) String() string {
	return fmt.Sprintf("%d", c)
}
