package models

import "errors"

type PlayingPiece interface {
	GetPiece() string
}

func GetPlayingPiece(pieceType string) (PlayingPiece, error) {
	switch pieceType {
	case "X":
		return GetPieceX(), nil
	case "O":
		return GetPieceO(), nil
	default:
		return nil, errors.New("invalid piece type")
	}
}
