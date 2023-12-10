package controllers

import (
	"fmt"
	"pattern/SnakeLadders/V1/models"
)

type GameContoller struct {
	game *models.Game
}

func GetGameContoller() *GameContoller {
	return &GameContoller{}
}

func (g *GameContoller) LaunchGame() {
	g.game = models.InitialiseGame(nil)
}

func (g *GameContoller) AddPlayer(player *models.Player) {
	g.game.AddPlayer(player)
}

func (g *GameContoller) MakeMove(player *models.Player, endCell *models.Cell) bool {
	if g.game.GetBoard().IsSnake(endCell) {
		endCell = g.game.GetBoard().GetSnakeDest(endCell)
	}

	if g.game.GetBoard().IsLadder(endCell) {
		endCell = g.game.GetBoard().GetLadderDest(endCell)
	}

	if IsWinner(endCell) {
		fmt.Println("Player ", player.GetName(), " to ", endCell.String())
		return true
	}
	if player.Move(endCell) {
		fmt.Println("Player ", player.GetName(), " to ", endCell.String())
	} else {
		fmt.Println("Player ", player.GetName(), " couldnot make a move")
	}
	return false
}

func (g *GameContoller) GetCurrPlayerTurn() *models.Player {
	return g.game.GetCurrPlayerTurn()
}

func (g *GameContoller) RollDice() int {
	curVal := g.game.RollDice()
	for curVal%6 == 0 {
		curVal += g.game.RollDice()
	}
	return curVal
}

func IsWinner(cell *models.Cell) bool {
	return cell.GetCellNumber() == 100
}
