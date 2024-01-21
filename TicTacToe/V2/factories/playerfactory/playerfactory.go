package playerfactory

import "pattern/TicTacToe/V2/classes"

type PlayerFactory struct{}

func (f *PlayerFactory) CreateBot(symbol classes.Symbol) classes.IPlayer {
	return classes.GetBotPlayer(symbol)
}

func (f *PlayerFactory) CreateBotWithStrategy(symbol classes.Symbol, botPlayingStrategy classes.BotPlayingStrategy) classes.IPlayer {
	return classes.GetBotPlayer(symbol).WithStrategy(botPlayingStrategy)
}

func (f *PlayerFactory) CreateHuman(symbol classes.Symbol) classes.IPlayer {
	return classes.GetHumanPlayer(symbol)
}
