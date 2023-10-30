package entities

type ExpenseOwingUser struct {
	IExpense
	IUser
	amount float64
}

func NewExpenseOwingUser(expense IExpense, user IUser, amount float64) *ExpenseOwingUser {
	return &ExpenseOwingUser{
		IExpense: expense,
		IUser:    user,
		amount:   amount,
	}
}

func (e *ExpenseOwingUser) GetAmount() float64 {
	return e.amount
}
