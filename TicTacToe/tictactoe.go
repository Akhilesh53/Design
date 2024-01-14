package tictactoe

import (
	"fmt"
	"pattern/TicTacToe/V1/models"
)

//for v1
func CallTicTacToe() {
	game, err := models.InitialiseGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	game.StartGame()
}

