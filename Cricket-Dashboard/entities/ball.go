package entities

import "pattern/Cricket-Dashboard/enums"

type Ball struct {
	ballNumber uint16
	ballType   enums.BallType
	bowledBy   Player
	facedBy    Player
	runType    enums.RunType
	wicket     Wicket
	// commentary
}

func NewBall(ballNumber uint16) Ball{
	return Ball{ballNumber: ballNumber}
}

func (b Ball) OfBallType(ballType enums.BallType) Ball{
	b.ballType = ballType
	return b
}

func (b Ball) GivingRun(runs enums.RunType) Ball{
	b.runType = runs
	return b
}