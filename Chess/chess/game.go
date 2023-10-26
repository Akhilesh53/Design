package chess

type Game struct {
	board   *Board
	player1 IPlayer
	player2 IPlayer
}

func NewGame(player1, player2 IPlayer) *Game {
	game := &Game{
		board:   InitialiseBoard(8),
		player1: player1,
		player2: player2,
	}
	return game
}

func (g *Game) GetBoard() *Board {
	return g.board
}

func (g *Game) GetPlayer1() IPlayer {
	return g.player1
}

func (g *Game) GetPlayer2() IPlayer {
	return g.player2
}

func (g *Game) StartGame() {

	var move IMove

	// for an infinite loop
	for {
		// ask player 1 which piece he wants to move - get piece
		// check if any of the selected piece is alive for this player
		var piece IPiece

		// if yes
		// ask player from which spot/ position - get startspot
		// check if start spot has the same piece or not
		var start ISpot

		//if yes
		// ask player 1 where he wants to move the piece - get endspot
		var end ISpot

		// check if the piece can move to endspot or not
		if !piece.CanMove(start, end) {
			// print cannot move
			continue //with the same player turn
		}

		// if everthing is fine
		move = NewMove(start, end, g.player1)
		move.MakeMove()
	}
}
