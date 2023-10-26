package chess

type IPlayer interface {
	GetPlayerName() string
	SetPlayerName(name string)
	IsWhiteSide() bool
	SetWhiteSide(isWhiteSide bool)
	IsHuman() bool
	SetHuman(isHuman bool)
}

func NewPlayer(name string, isWhiteSide bool, isHuman bool) IPlayer {
	return &player{
		name:        name,
		isWhiteSide: isWhiteSide,
		isHuman:     isHuman,
	}
}

type player struct {
	name        string
	isWhiteSide bool
	isHuman     bool
}

func (p *player) SetPlayerName(name string) {
	p.name = name
}

func (p *player) SetWhiteSide(isWhiteSide bool) {
	p.isWhiteSide = isWhiteSide
}

func (p *player) SetHuman(isHuman bool) {
	p.isHuman = isHuman
}

func (p *player) GetPlayerName() string {
	return p.name
}

func (p *player) IsWhiteSide() bool {
	return p.isWhiteSide
}

func (p *player) IsHuman() bool {
	return p.isHuman
}
