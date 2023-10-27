package entities

type IAccountNumberGenerator interface {
	Generate() int
}

type accountNumberGenerator struct {
	number int
}

func NewAccountNumberGenerator() IAccountNumberGenerator {
	return &accountNumberGenerator{
		number: 0,
	}
}

func (w *accountNumberGenerator) Generate() int {
	w.number++
	return w.number
}
