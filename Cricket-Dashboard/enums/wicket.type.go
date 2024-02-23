package enums

type WicketType int

const (
	BOWLED WicketType = iota
	CAUGHT
	RUN_OUT
	HIT_WICKET
	STUMPTED
	LBW
	OBSTRUCTING
	RETIRED_HURT
)
