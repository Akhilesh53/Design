package main

import "fmt"

func main() {
	left := NewNumber(5)
	right := NewNumber(10)

	exp1 := NewExpression(left, right, "+")
	fmt.Println("sum: ", exp1.Evaluate())

	left1 := NewNumber(50)
	exp2 := NewExpression(left1, exp1, "*") //50*(5+10)

	fmt.Println("exp2: ", exp2.Evaluate())
}
