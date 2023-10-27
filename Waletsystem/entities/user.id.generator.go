package entities

type IUserIdGenerator interface {
	Generate() int
}

type userIdGenerator struct {
	number int
}

func NewUserIdGenerator() IUserIdGenerator {
	return &userIdGenerator{
		number: 0,
	}
}

func (u *userIdGenerator) Generate() int {
	u.number++
	return u.number
}
