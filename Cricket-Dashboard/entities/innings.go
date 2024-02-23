package entities

type Innings struct {
	inningNumber uint16
	playerScores map[string]PlayerScore
	bowlerOvers  map[uint16]Over
	totalExtras  uint16
	totalScores  uint16
	totalWickets uint16
	totalByes    uint16
	totalLegByes uint16
	mactchId     uint16
}

func NewInnings(inningsNumber, matchId uint16) Innings {
	return Innings{
		inningNumber: inningsNumber,
		mactchId:     matchId,
		playerScores: make(map[string]PlayerScore),
		bowlerOvers:  make(map[uint16]Over),
	}
}

func (i Innings) AddOver(over Over) {
	i.bowlerOvers[over.overNumber] = over
}
