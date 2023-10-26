package chess

import "math"

// define an interface for chess borad piece
type IPiece interface {
	GetPieceType() PieceType
	SetPieceType(pieceType PieceType)
	IsWhite() bool
	SetWhite(isWhite bool)
	IsAlive() bool
	SetAlive(alive bool)
	CanMove(startSpot, endSpot ISpot) bool
}

func NewPiece(pieceType PieceType, isWhite bool) IPiece {
	switch pieceType {
	case PAWN:
		return newPawn(isWhite)
	case ROOK:
		return newRook(isWhite)
	case KNIGHT:
		return newKnight(isWhite)
	case BISHOP:
		return newBishop(isWhite)
	case QUEEN:
		return newQueen(isWhite)
	case KING:
		return newKing(isWhite)
	default:
		return defaultpiece()
	}
}

func defaultpiece() IPiece {
	return &defaultPiece{
		pieceType: UNKNOWN_PIECETYPE,
		isWhite:   false,
		isAlive:   false,
	}
}

func (p *defaultPiece) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *defaultPiece) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *defaultPiece) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *defaultPiece) CanMove(startSpot, endSpot ISpot) bool {
	return false
}

func (p *defaultPiece) GetPieceType() PieceType {
	return p.pieceType
}

func (p *defaultPiece) IsWhite() bool {
	return p.isWhite
}

func (p *defaultPiece) IsAlive() bool {
	return p.isAlive
}

type defaultPiece struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

type king struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newKing(isWhite bool) IPiece {
	return &king{
		pieceType: KING,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *king) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *king) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *king) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *king) CanMove(startSpot, endSpot ISpot) bool {
	//check whether king can move to endspot or not
	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *king) GetPieceType() PieceType {
	return p.pieceType
}

func (p *king) IsWhite() bool {
	return p.isWhite
}

func (p *king) IsAlive() bool {
	return p.isAlive
}

type queen struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newQueen(isWhite bool) IPiece {
	return &queen{
		pieceType: QUEEN,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *queen) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *queen) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *queen) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *queen) CanMove(startSpot, endSpot ISpot) bool {
	//check whether queen can move to endspot or not
	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *queen) GetPieceType() PieceType {
	return p.pieceType
}

func (p *queen) IsWhite() bool {
	return p.isWhite
}

func (p *queen) IsAlive() bool {
	return p.isAlive
}

type bishop struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newBishop(isWhite bool) IPiece {
	return &bishop{
		pieceType: BISHOP,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *bishop) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *bishop) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *bishop) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *bishop) CanMove(startSpot, endSpot ISpot) bool {
	//check whether bishop can move to endspot or not

	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *bishop) GetPieceType() PieceType {
	return p.pieceType
}

func (p *bishop) IsWhite() bool {
	return p.isWhite
}

func (p *bishop) IsAlive() bool {
	return p.isAlive
}

type knight struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newKnight(isWhite bool) IPiece {
	return &knight{
		pieceType: KNIGHT,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *knight) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *knight) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *knight) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *knight) CanMove(startSpot, endSpot ISpot) bool {
	//check whether knight can move to endspot or not
	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *knight) GetPieceType() PieceType {
	return p.pieceType
}

func (p *knight) IsWhite() bool {
	return p.isWhite
}

func (p *knight) IsAlive() bool {
	return p.isAlive
}

type rook struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newRook(isWhite bool) IPiece {
	return &rook{
		pieceType: ROOK,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *rook) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *rook) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *rook) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *rook) CanMove(startSpot, endSpot ISpot) bool {
	//check whether rook can move to endspot or not
	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *rook) GetPieceType() PieceType {
	return p.pieceType
}

func (p *rook) IsWhite() bool {
	return p.isWhite
}

func (p *rook) IsAlive() bool {
	return p.isAlive
}

type pawn struct {
	pieceType PieceType
	isWhite   bool
	isAlive   bool
}

func newPawn(isWhite bool) IPiece {
	return &pawn{
		pieceType: PAWN,
		isWhite:   isWhite,
		isAlive:   true,
	}
}

func (p *pawn) SetPieceType(pieceType PieceType) {
	p.pieceType = pieceType
}

func (p *pawn) SetWhite(isWhite bool) {
	p.isWhite = isWhite
}

func (p *pawn) SetAlive(alive bool) {
	p.isAlive = alive
}

func (p *pawn) CanMove(startSpot, endSpot ISpot) bool {
	//check whether pawn can move to endspot or not
	if endSpot.GetPiece().IsWhite() == startSpot.GetPiece().IsWhite() {
		return false
	}

	return math.Abs(float64(startSpot.GetPosition().GetX())-float64(endSpot.GetPosition().GetX())) <= 1 && math.Abs(float64(startSpot.GetPosition().GetY())-float64(endSpot.GetPosition().GetY())) <= 1
}

func (p *pawn) GetPieceType() PieceType {
	return p.pieceType
}

func (p *pawn) IsWhite() bool {
	return p.isWhite
}

func (p *pawn) IsAlive() bool {
	return p.isAlive
}
