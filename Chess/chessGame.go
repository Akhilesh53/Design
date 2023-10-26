package main

import "pattern/Chess/chess"

func main() {

	player1 := chess.NewPlayer("Player 1", true, true)
	player2 := chess.NewPlayer("Player 2", false, false)

	game := chess.NewGame(player1, player2)

	game.GetBoard().AddSpot(chess.NewSpot(chess.NewPiece(chess.PAWN, true), chess.NewPosition(1, 0)))

	game.GetBoard().PrintBoard()
}
