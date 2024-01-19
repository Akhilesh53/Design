package classes

type HumanPlayer struct {
	playerType PlayerType
	symbol     Symbol
}

func (player *HumanPlayer) GetPlayerType() PlayerType {
	return player.playerType
}

func (player *HumanPlayer) GetPlayerSymbol() Symbol {
	return player.symbol
}

func (player *HumanPlayer) getPlayer(symbol Symbol) IPlayer {
	return &HumanPlayer{
		playerType: HUMAN,
		symbol:     symbol,
	}
}

func GetHumanPlayer(symbol Symbol) IPlayer {
	h := HumanPlayer{}
	return h.getPlayer(symbol)
}
