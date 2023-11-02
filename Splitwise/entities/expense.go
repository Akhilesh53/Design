package entities

import "time"

type IExpense interface {
	GetId() int
	GetDescription() string
	GetAmount() float64
	GetCreatedAt() time.Time
	GetCreatedBy() IUser
	GetExpenseCurrency() Currency
	GetExpenseParticipants() []IUser
	GetLastModified() time.Time
	SetId(int) IExpense
	SetDescription(string) IExpense
	SetAmount(float64) IExpense
	SetCreatedAt(time.Time) IExpense
	SetCreatedBy(IUser) IExpense
	SetExpenseCurrency(Currency) IExpense
	SetExpenseParticipants(users []IUser) IExpense
	AddExpenseParticipant(user IUser) IExpense
	RemoveExpenseParticipant(user IUser) IExpense
	SetLastModified(time.Time) IExpense
}

type expense struct {
	id           int
	description  string
	amount       float64
	createdAt    time.Time
	createdBy    IUser
	currency     Currency
	participants []IUser
	lastModified time.Time
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

func (e *expense) GetExpenseCurrency() Currency {
	return e.currency
}

func (e *expense) SetExpenseCurrency(currency Currency) IExpense {
	e.currency = currency
	return e
}

func (e *expense) GetExpenseParticipants() []IUser {
	return e.participants
}

func (e *expense) SetExpenseParticipants(users []IUser) IExpense {
	e.participants = users
	return e
}

func (e *expense) AddExpenseParticipant(user IUser) IExpense {
	e.participants = append(e.participants, user)
	return e
}

func (e *expense) RemoveExpenseParticipant(user IUser) IExpense {
	for i, participant := range e.participants {
		if participant.GetId() == user.GetId() {
			e.participants = append(e.participants[:i], e.participants[i+1:]...)
			break
		}
	}
	return e
}

func (e *expense) GetLastModified() time.Time {
	return e.lastModified
}

func (e *expense) SetLastModified(lastModified time.Time) IExpense {
	e.lastModified = lastModified
	return e
}
