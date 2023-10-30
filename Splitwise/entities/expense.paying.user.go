package entities

type ExpensePayingUser struct {
	IExpense
	IUser
	amount float64
}

func NewExpensePayingUser(expense IExpense, user IUser, amount float64) *ExpensePayingUser {
	return &ExpensePayingUser{
		IExpense: expense,
		IUser:    user,
		amount:   amount,
	}
}

func (e *ExpensePayingUser) GetAmount() float64 {
	return e.amount
}
