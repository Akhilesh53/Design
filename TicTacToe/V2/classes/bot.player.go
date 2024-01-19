package classes

type BotPlayer struct {
	playerType PlayerType
	symbol     Symbol
}

func (player *BotPlayer) GetPlayerType() PlayerType {
	return player.playerType
}

func (player *BotPlayer) GetPlayerSymbol() Symbol {
	return player.symbol
}

func (player *BotPlayer) getPlayer(symbol Symbol) IPlayer {
	return &BotPlayer{
		playerType: HUMAN,
		symbol:     symbol,
	}
}

func GetBotPlayer(symbol Symbol) IPlayer {
	h := BotPlayer{}
	return h.getPlayer(symbol)
}
