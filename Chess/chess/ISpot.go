package chess

// define an interface ISpot whihc is implemented by Spot and has piece and position
type ISpot interface {
	GetPiece() IPiece
	SetPiece(piece IPiece) ISpot
	GetPosition() Position
	SetPosition(position Position) ISpot
	HasThisPiece(piece IPiece) bool
	HasPiece() bool
}

func NewSpot(piece IPiece, position Position) ISpot {
	return &spot{
		piece:    piece,
		position: position,
	}
}

type spot struct {
	piece    IPiece
	position Position
}

func (s *spot) SetPiece(piece IPiece) ISpot {
	s.piece = piece
	return s
}

func (s *spot) SetPosition(position Position) ISpot {
	s.position = position
	return s
}

func (s *spot) HasThisPiece(piece IPiece) bool {
	return s.piece == piece
}

func (s *spot) HasPiece() bool {
	return s.piece != nil
}

func (s *spot) GetPiece() IPiece {
	return s.piece
}

func (s *spot) GetPosition() Position {
	return s.position
}

type Position struct {
	x int
	y int
}

func NewPosition(x, y int) Position {
	return Position{x: x, y: y}
}

func (p Position) GetX() int {
	return p.x
}

func (p Position) GetY() int {
	return p.y
}
