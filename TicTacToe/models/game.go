package models

import "fmt"

type Game struct {
	Players []*Player
	Board   *Board
}

func InitialiseGame() *Game {
	players := []*Player{}
	piecex, err := GetPlayingPiece("X")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	piecey, err := GetPlayingPiece("O")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	players = append(players, GetPlayer("Player1", piecex))
	players = append(players, GetPlayer("Player2", piecey))
	return &Game{Players: players, Board: InitialiseBoard(3)}
}

func (g *Game) StartGame() {

	gameContinue := true
	for gameContinue {

        // get the count of vacant places
		// if none of the place left and no player won, means game tie
		if g.Board.GetVacantPlaces() == 0{
			fmt.Println(" Ohh !!! \n Game Tied....")
			gameContinue = false
			continue
		}

		g.Board.PrintBoard()
		playerTurn := g.Players[0]

		var row int
		var col int
		fmt.Println("Enter row, col where you want to insert piece :: ", playerTurn.GetName())
		_, err := fmt.Scanln(&row, &col)
		if err != nil {
			fmt.Println(err)
			return
		}

		// check if some value is alredy present or not at req position
		// if not then only it will insert
		if !g.Board.IsNull(row, col) {
			fmt.Println("You " + playerTurn.GetName() + " have chosen wrong position, Please choose again.")
			continue
		}

		g.Board.AddPiece(playerTurn.GetPiece(), row, col)

		if g.IsPlayerWinner(playerTurn) {
			fmt.Println("Congratulations !!! \n" + playerTurn.GetName() + " has won the game ....")
			gameContinue = false
			continue
		}

		g.Players = g.Players[1:]
		g.Players = append(g.Players, playerTurn)

	}
}

func (g *Game) IsPlayerWinner(player *Player) bool {

	// check if the current player has same piece in any row, column or diagonal
	return g.Board.IsPieceInAnyRow(player.GetPiece()) || g.Board.IsPieceInAnyColumn(player.GetPiece()) || g.Board.IsPieceInAnyDiagonal(player.GetPiece())

}
