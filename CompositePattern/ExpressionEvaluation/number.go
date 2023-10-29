package main

import "fmt"

type Number struct {
	value int
}

func NewNumber(value int) IExpression {
	return &Number{
		value: value,
	}
}

func (n *Number) Evaluate() int{
	fmt.Println("Number: ", n.value)
	return n.value
}
