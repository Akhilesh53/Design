package main

type expression struct {
	left      IExpression
	right     IExpression
	operation string
}

func NewExpression(left IExpression, right IExpression, operation string) IExpression {
	return &expression{
		left:      left,
		right:     right,
		operation: operation,
	}
}

func (e *expression) Evaluate() int {
	var ans int
	switch e.operation {
	case "+":
		ans = e.left.Evaluate() + e.right.Evaluate()
	case "-":
		ans = e.left.Evaluate() - e.right.Evaluate()
	case "*":
		ans = e.left.Evaluate() * e.right.Evaluate()
	case "/":
		ans = e.left.Evaluate() / e.right.Evaluate()
	default:
		ans = 0
	}
	return ans
}
