package models

type Game struct {
	board      *Board
	players    []*Player
	dice       *Dice
	currPlayer int
	maxPlayers int
}

func InitialiseGame(players []*Player) *Game {
	if players == nil {
		players = make([]*Player, 0)
	}
	return &Game{
		board:      InitializeBoard(),
		players:    players,
		dice:       NewDice(),
		currPlayer: 0,
		maxPlayers: len(players),
	}
}

func (g *Game) GetBoard() *Board {
	return g.board
}

func (g *Game) GetPlayers() []*Player {
	return g.players
}

func (g *Game) GetDice() *Dice {
	return g.dice
}

func (g *Game) SetBoard(board *Board) *Game {
	g.board = board
	return g
}

func (g *Game) SetPlayers(players []*Player) *Game {
	g.players = players
	g.maxPlayers = len(players)
	return g
}

func (g *Game) SetDice(dice *Dice) *Game {
	g.dice = dice
	return g
}

func (g *Game) AddPlayer(player *Player) *Game {
	g.players = append(g.players, player)
	g.maxPlayers = len(g.players)
	return g
}

func (g *Game) GetCurrPlayerTurn() *Player {
	currPlayer := g.currPlayer
	g.currPlayer++
	if g.currPlayer == g.maxPlayers {
		g.currPlayer = 0
	}
	return g.players[currPlayer]
}

func (g *Game) RollDice() int {
	return g.dice.Roll()
}
