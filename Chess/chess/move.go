package chess

// there are two ways to defina a move
//1) either we define a makeMove method in Player interface
//2) or we define a Move struct and pass it to the Game object

type IMove interface {
	GetStartSpot() ISpot
	SetStartSpot(spot ISpot)
	GetEndSpot() ISpot
	SetEndSpot(spot ISpot)
	GetPlayer() IPlayer
	SetPlayer(player IPlayer)
	MakeMove()
}

type move struct {
	startSpot ISpot
	endSpot   ISpot
	player    IPlayer
}

func NewMove(startSpot, endSpot ISpot, player IPlayer) IMove {
	return &move{
		startSpot: startSpot,
		endSpot:   endSpot,
		player:    player,
	}
}

func (m *move) SetStartSpot(spot ISpot) {
	m.startSpot = spot
}

func (m *move) GetStartSpot() ISpot {
	return m.startSpot
}

func (m *move) GetEndSpot() ISpot {
	return m.endSpot
}

func (m *move) SetEndSpot(spot ISpot) {
	m.endSpot = spot
}

func (m *move) GetPlayer() IPlayer {
	return m.player
}

func (m *move) SetPlayer(player IPlayer) {
	m.player = player
}

func (m *move) MakeMove() {
	// check if endspot has piece of opposite colour
	if m.GetEndSpot().HasPiece() && m.GetEndSpot().GetPiece().IsWhite() != m.GetStartSpot().GetPiece().IsWhite() {
		m.GetEndSpot().GetPiece().SetAlive(false)
	}
	// if yes, kill it
	// else move the piece to endspot
	// change the spots of the piece
	m.SetEndSpot(m.GetEndSpot().SetPiece(m.GetStartSpot().GetPiece()))
	m.SetStartSpot(m.GetStartSpot().SetPiece(defaultpiece()))
}
