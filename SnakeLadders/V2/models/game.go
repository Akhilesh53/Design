package models

import (
	"errors"
	"fmt"
)

type Game struct {
	board            *Board
	players          []*Player
	dice             *Dice
	currPlayer       int
	maxPlayers       int
	colourAssignment map[Color]bool
}

func InitialiseGame(players []*Player) (*Game, error) {
	board := InitializeBoard()
	if players == nil {
		players = make([]*Player, 0)
	}

	if len(players) > 4 {
		return nil, errors.New("only 4 players can play a game at a time")
	}

	colormap := make(map[Color]bool)

	for _, p := range players {
		if colormap[p.GetColor()] {
			return nil, errors.New("colour is already assigned")
		}
		colormap[p.GetColor()] = true
		p.SetPosition(board.GetCell(9, 0))
	}
	return &Game{
		board:            board,
		players:          players,
		dice:             NewDice(),
		currPlayer:       0,
		maxPlayers:       len(players),
		colourAssignment: colormap,
	}, nil
}

func (g *Game) GetCurrPlayerTurn() *Player {
	if g.currPlayer == g.maxPlayers {
		g.currPlayer = 0
	}

	player := g.players[g.currPlayer]
	g.currPlayer = (g.currPlayer + 1) % g.maxPlayers
	return player
}

func (g *Game) RollDice() int {
	curr6count := 0
	currVal := g.dice.Roll()
	for currVal%6 == 0 {
		if curr6count == 3 {
			fmt.Println("your trun got cancelled, three times 6 occured.")
			return 0
		}
		curr6count++
		val := g.dice.Roll()
		currVal += val
	}
	fmt.Println("Dice Value :: ", currVal)
	return currVal
}

func (g *Game) GetBoard() *Board {
	return g.board
}

func (g *Game) PrintBoard() {
	g.board.Print()
}

func (g *Game) GetCellByCellNumber(cellNumber int) *Cell {
	return g.board.GetCellByCellNumber(cellNumber)
}