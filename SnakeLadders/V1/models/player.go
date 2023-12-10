package models

type Player struct {
	name     string
	color    Color
	position *Cell
}

func NewPlayer(name string, color Color) *Player {
	return &Player{name: name, color: color, position: NewCell(0)}
}

func NewDefaultPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) SetName(name string) *Player {
	p.name = name
	return p
}

func (p *Player) SetColor(color Color) *Player {
	p.color = color
	return p
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetColor() Color {
	return p.color
}

func (p *Player) GetPosition() *Cell {
	return p.position
}

func (p *Player) SetPosition(cell *Cell) *Player {
	p.position = cell
	return p
}

func (p *Player) Move(end *Cell) bool {
	if p.canMove(end) {
		p.SetPosition(end)
		return true
	}
	return false
}

func (p *Player) canMove(end *Cell) bool {
	return end.GetCellNumber() <= 100
}
