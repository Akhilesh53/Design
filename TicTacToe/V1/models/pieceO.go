package models

type PieceO struct {
	pieceName string
}

func GetPieceO() *PieceO {
	return &PieceO{
		pieceName: "O",
	}
}

func (x *PieceO) GetPiece() string {
	return x.pieceName
}
