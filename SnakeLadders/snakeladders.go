package main

import (
	"fmt"
	"pattern/SnakeLadders/V2/controllers"
	"pattern/SnakeLadders/V2/models"
)

func main() {
	gc := controllers.GetGameContoller()

	players := make([]*models.Player, 0)
	players = append(players, models.NewPlayer("Akhilesh", models.RED))
	players = append(players, models.NewPlayer("Mahajan", models.BLUE))
	if err := gc.LaunchGame(players); err != nil {
		fmt.Println(err)
		return
	}

	gc.PrintBoard()
	gameContinue := true

	for gameContinue {
		currPlayer := gc.GetCurrPlayerTurn()

		fmt.Println("Enter `r` to roll the dice for player turn ", currPlayer.GetName(), currPlayer.GetPosition(), currPlayer.GetPosition().GetCellNumber())
		var inp string
		if _, err := fmt.Scan(&inp); err != nil {
			fmt.Println(err)
			break
		}
		diceVal := gc.RollDice()

		start := currPlayer.GetPosition().GetCellNumber()
		fmt.Println(start + diceVal)
		targetCell := gc.GetCellByCellNumber(start + diceVal)
		currPlayer.SetPosition(targetCell)
		fmt.Println(currPlayer.GetPosition().GetCellNumber())
		if gc.IsWinningMove(currPlayer) {
			fmt.Println(currPlayer.GetName(), "wins the game")
			gameContinue = false
		}
	}
}
