package models

type PieceType int

const (
	X PieceType = iota
	O
)

func (p PieceType) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return ""
	}
}