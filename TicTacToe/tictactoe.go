package tictactoe

import (
	"pattern/TicTacToe/models"
)

func CallTicTacToe() {
	game := models.InitialiseGame()
	game.StartGame()
}
