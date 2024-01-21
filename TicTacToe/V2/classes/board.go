package classes

type Board struct {
	dimension int
	board     [][]Cell
}

func NewBoardWithDimension(dimension int) *Board {
	board := make([][]Cell, dimension)
	for i := 0; i < dimension; i++ {
		board[i] = make([]Cell, dimension)
	}
	return &Board{dimension: dimension, board: board}
}
