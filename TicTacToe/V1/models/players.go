package models

type Player struct {
	name  string
	piece PlayingPiece
}

func GetPlayer(name string, piece PlayingPiece) *Player {
	return &Player{
		name:  name,
		piece: piece,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPiece() PlayingPiece {
	return p.piece
}
