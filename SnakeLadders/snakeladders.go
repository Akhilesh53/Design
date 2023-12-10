package main

import (
	"fmt"
	"pattern/SnakeLadders/V1/controllers"
	"pattern/SnakeLadders/V1/models"
)

func main() {
	gc := controllers.GetGameContoller()

	gc.LaunchGame()

	//Add players to the game
	gc.AddPlayer(models.NewPlayer("Akhilesh", models.RED))
	gc.AddPlayer(models.NewPlayer("Mahajan", models.BLUE))

	gameContinue := true

	for gameContinue {
		currPlayer := gc.GetCurrPlayerTurn()

		fmt.Println("Enter `r` to roll the dice")
		var inp string
		if _, err := fmt.Scan(&inp); err != nil {
			fmt.Println(err)
			break
		}
		diceVal := gc.RollDice()

		start := currPlayer.GetPosition()
		start.SetCellNumber(start.GetCellNumber() + diceVal)

		if gc.MakeMove(currPlayer, start) {
			fmt.Println(currPlayer.GetName(), "wins the game")
			gameContinue = false
		}
	}
}
