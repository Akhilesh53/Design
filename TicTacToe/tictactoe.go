package tictactoe

import (
	"fmt"
	"pattern/TicTacToe/models"
)

func CallTicTacToe() {
	game, err := models.InitialiseGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	game.StartGame()
}
