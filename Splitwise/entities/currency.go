package entities

type Currency int

const (
	USD Currency = iota
	INR
	EUR
	GBP
)

func (c Currency) String() string {
	switch c {
	case USD:
		return "USD"
	case INR:
		return "INR"
	case EUR:
		return "EUR"
	case GBP:
		return "GBP"
	default:
		return "UNKNOWN"
	}
}
