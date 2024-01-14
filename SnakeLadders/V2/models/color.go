package models

import "fmt"

type Color int

const (
	RED Color = iota
	BLUE
	YELLOW
	GREEN
)

var ColorsArray = []Color{RED, BLUE, YELLOW, GREEN}

func (c *Color) String() string {
	return fmt.Sprintf("%d", c)
}
