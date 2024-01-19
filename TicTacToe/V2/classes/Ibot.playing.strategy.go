package classes

type BotPlayingStratrgy interface {
	MakeNextMove(board Board, symbol Symbol) Move
}
