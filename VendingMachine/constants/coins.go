package constants

type Coins int

const (
	Penny   Coins = 1
	Nickel  Coins = 5
	Dime    Coins = 10
	Quarter Coins = 25
)

func (c Coins) Value() int {
	switch c {
	case Penny:
		return 1
	case Nickel:
		return 5
	case Dime:
		return 10
	case Quarter:
		return 25
	default:
		return 0
	}
}
