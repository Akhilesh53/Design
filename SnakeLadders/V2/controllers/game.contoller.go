package controllers

import (
	"fmt"
	"pattern/SnakeLadders/V2/models"
)

type GameContoller struct {
	game *models.Game
}

func GetGameContoller() *GameContoller {
	return &GameContoller{}
}

func (g *GameContoller) LaunchGame(players []*models.Player) error {
	var err error
	g.game, err = models.InitialiseGame(players)
	return err
}

func (g *GameContoller) IsWinningMove(player *models.Player) bool {
	if player.GetPosition().IsSnake() {
		dest := g.game.GetBoard().GetSnakesMap().GetTail(player.GetPosition().GetCellNumber())
		fmt.Println("!!!!!! Snake bite :( \nYour new destination is ", dest)
		player.SetPosition(g.GetCellByCellNumber(dest))
	}

	if player.GetPosition().IsLadder() {
		dest := g.game.GetBoard().GetLaddersMap().GetTop(player.GetPosition().GetCellNumber())
		fmt.Println("Wuhuuuuu..... A ladder \nYour new destination is ", dest)
		player.SetPosition(g.GetCellByCellNumber(dest))
	}

	if IsWinner(player.GetPosition()) {
		fmt.Println("Player ", player.GetName(), " to ", player.GetPosition().String())
		return true
	}
	fmt.Println("Player ", player.GetName(), " to ", player.GetPosition().GetCellNumber())
	return false
}

func (g *GameContoller) GetCurrPlayerTurn() *models.Player {
	return g.game.GetCurrPlayerTurn()
}

func (g *GameContoller) RollDice() int {
	return g.game.RollDice()
}

func IsWinner(cell *models.Cell) bool {
	return cell.GetCellNumber() == 100
}

func (g *GameContoller) PrintBoard() {
	g.game.PrintBoard()
}

func (g *GameContoller) GetCellByCellNumber(cell int) *models.Cell {
	return g.game.GetCellByCellNumber(cell)
}
