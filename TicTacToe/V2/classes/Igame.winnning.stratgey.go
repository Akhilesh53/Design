package classes


type GameWinningStrategy interface {
	CheckIfWon(board Board, lastMovedCell Cell) bool
}

