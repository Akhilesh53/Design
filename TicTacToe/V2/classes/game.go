package classes

import "pattern/TicTacToe/V2/errors"

type Game struct {
	board                *Board
	players              []IPlayer
	moves                []*Move
	winningStrategies    []GameWinningStrategy
	lastMoverPlayerIndex int
	status               GameStatus
	winner               IPlayer
}

// todo:
// to build a game, we need to build / set attributes for different classes.
// that can be easily done by builder design pattern
func (game *Game) GetGameStatus() GameStatus {
	return game.status
}

// define a builder object with the same attributes as concrete class
type GameBuilder struct {
	game Game
}

func NewGameBuilder() *GameBuilder {
	return &GameBuilder{}
}

// define setter methods for all attributes of class
// all the setter methods will return the builder object itself
func (gb *GameBuilder) WithBoard(board *Board) *GameBuilder {
	gb.game.board = board
	return gb
}

func (gb *GameBuilder) SetStatus(status GameStatus) *GameBuilder {
	gb.game.status = status
	return gb
}

// rather than having list of players in attributes it is better to add player one by one
// it will make the code more readible
func (gb *GameBuilder) AddPlayer(player IPlayer) *GameBuilder {
	gb.game.players = append(gb.game.players, player)
	return gb
}

//todo:
// internal state params should not be allowed to set by builder class
// like status winner, lastPlayerIndex

// there should be one method for builder class which will prepare the response
// it shoould be taken care while building the game all validatons should be taken care of
// 1) there should be only 1 bot player
// 2) there should be no teo players with same symbol
func (gb *GameBuilder) Build() (*Game, error) {
	if !gb.checkIfOneBotMax() {
		return nil, errors.ErrMultipleBots
	}
	return &gb.game, nil
}

func (gb *GameBuilder) checkIfOneBotMax() bool {
	var cnt = 0

	for _, player := range gb.game.players {
		if player.GetPlayerType() == BOT {
			cnt++
		}
	}
	return cnt <= 1
}
