package tictactoe

import (
	"fmt"
	"pattern/TicTacToe/V1/models"
	"pattern/TicTacToe/V2/controllers"
)

// for v1
func CallTicTacToe() {
	game, err := models.InitialiseGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	game.StartGame()
}

// for v2

func TicTacToeApplicationCode() {

	gameController := controllers.MakeGameContoller()
	// take necessary params while creating game like dimension, list of players
	game, err := gameController.CreateGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(game)
}
