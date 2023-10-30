package entities

import "time"

type IExpense interface {
	GetId() int
	GetDescription() string
	GetAmount() float64
	GetCreatedAt() time.Time
	GetCreatedBy() IUser
	SetId(int) IExpense
	SetDescription(string) IExpense
	SetAmount(float64) IExpense
	SetCreatedAt(time.Time) IExpense
	SetCreatedBy(IUser) IExpense
}

type expense struct {
	id          int
	description string
	amount      float64
	createdAt   time.Time
	createdBy   IUser
}

func NewExpense(id int, description string, amount float64, createdBy IUser) IExpense {
	return &expense{
		id:          id,
		description: description,
		amount:      amount,
		createdAt:   time.Now(),
		createdBy:   createdBy,
	}
}

func (e *expense) GetId() int {
	return e.id
}

func (e *expense) GetDescription() string {
	return e.description
}

func (e *expense) GetAmount() float64 {
	return e.amount
}

func (e *expense) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e *expense) GetCreatedBy() IUser {
	return e.createdBy
}

func (e *expense) SetId(id int) IExpense {
	e.id = id
	return e
}

func (e *expense) SetDescription(description string) IExpense {
	e.description = description
	return e
}

func (e *expense) SetAmount(amount float64) IExpense {
	e.amount = amount
	return e
}

func (e *expense) SetCreatedAt(createdAt time.Time) IExpense {
	e.createdAt = createdAt
	return e
}

func (e *expense) SetCreatedBy(createdBy IUser) IExpense {
	e.createdBy = createdBy
	return e
}
