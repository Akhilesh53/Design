package classes

type BotPlayingStrategy interface {
	MakeNextMove(board Board, symbol Symbol) Move
}
