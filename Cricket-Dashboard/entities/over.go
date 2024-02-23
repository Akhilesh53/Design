package entities

type Over struct {
	overNumber uint16
	totalRuns  uint16
	extraScore uint16
	runsScored uint16
	balls      []Ball
}

func NewOver(overNumber uint16) Over {
	return Over{overNumber: uint16(overNumber)}
}

func (o Over) AddBall(ball Ball) Over {
	o.balls = append(o.balls, ball)
	return o
}
