package controllers

import "pattern/TicTacToe/V2/classes"

// all the actions in main class should not be done by object diectly
// all actions hould be done by contorller  only
type GameController struct {
}

func MakeGameContoller() *GameController {
	return &GameController{}
}

// actions that can be taken over a game
// create/ initialise game
// make a move
// undo
// check winner
// check state
func (gc *GameController) CreateGame() (*classes.Game,error) {
	return nil,nil
}

func (gc *GameController) MakeMove(game classes.Game, move classes.Move) error {
	return nil
}

func (gc *GameController) Undo(game classes.Game) error {
	return nil
}

func (gc *GameController) CheckGameStatus(game classes.Game) classes.GameStatus {
	return game.GetGameStatus()
}
