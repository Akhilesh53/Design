package classes

type HumanPlayer struct {
	User
	playerType PlayerType
	symbol     Symbol
}

func (player *HumanPlayer) GetPlayerType() PlayerType {
	return player.playerType
}

func (player *HumanPlayer) GetPlayerSymbol() Symbol {
	return player.symbol
}

func (player *HumanPlayer) getPlayer(symbol Symbol) *HumanPlayer {
	return &HumanPlayer{
		playerType: HUMAN,
		symbol:     symbol,
	}
}

func GetHumanPlayer(symbol Symbol) *HumanPlayer {
	h := HumanPlayer{}
	return h.getPlayer(symbol)
}
