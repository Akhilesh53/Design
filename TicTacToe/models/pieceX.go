package models

type PieceX struct {
	pieceName PieceType
}

func GetPieceX() *PieceX {
	return &PieceX{
		pieceName: X,
	}
}

func (x *PieceX) GetPiece() string {
	return x.pieceName.String()
}
