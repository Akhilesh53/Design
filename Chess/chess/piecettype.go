package chess

type PieceType int

const (
	PAWN PieceType = iota
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
	UNKNOWN_PIECETYPE = -1
)

func (p PieceType) String() string {
	switch p {
	case PAWN:
		return "Pawn"
	case ROOK:
		return "Rook"
	case KNIGHT:
		return "Knight"
	case BISHOP:
		return "Bishop"
	case QUEEN:
		return "Queen"
	case KING:
		return "King"
	default:
		return "Unknown"
	}
}
