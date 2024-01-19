package tictactoe

import (
	"fmt"
	"pattern/TicTacToe/V1/models"
	"pattern/TicTacToe/V2/classes"
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

	game, err := classes.NewGameBuilder().
		WithBoard(classes.NewBoard(3)).
		AddPlayer(classes.GetHumanPlayer(classes.CROSS)).
		Build()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(game)
}
