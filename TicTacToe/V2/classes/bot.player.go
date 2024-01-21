package classes

type BotPlayer struct {
	User
	playingStrategy BotPlayingStrategy
	playerType      PlayerType
	symbol          Symbol
}

func (player *BotPlayer) GetPlayerType() PlayerType {
	return player.playerType
}

func (player *BotPlayer) GetPlayerSymbol() Symbol {
	return player.symbol
}

func (player *BotPlayer) GetBotPlayingStrategy() BotPlayingStrategy {
	return player.playingStrategy
}
 
func (player *BotPlayer) getPlayer(symbol Symbol) *BotPlayer {
	return &BotPlayer{
		playerType: HUMAN,
		symbol:     symbol,
	}
}

func (player *BotPlayer) WithStrategy(strategy BotPlayingStrategy) *BotPlayer {
	player.playingStrategy = strategy
	return player
}

func GetBotPlayer(symbol Symbol) *BotPlayer {
	h := BotPlayer{}
	return h.getPlayer(symbol)
}
