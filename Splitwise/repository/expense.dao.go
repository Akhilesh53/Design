package repository

import (
	"errors"
	"pattern/Splitwise/entities"
)

type ExpenseDao interface {
	CreateExpense(id int, description string, amount float64, createdBy entities.IUser) (entities.IExpense, error)
	GetExpense(id int) (entities.IExpense, error)
	DeleteExpense(id int)
	UpdateExpense(id int, description string, amount float64, createdBy entities.IUser) (entities.IExpense, error)
	UpdateExpenseByParam(paramName, val interface{}, id int) (entities.IExpense, error)
}

type expenseDao struct {
	expenseDao map[int]entities.IExpense
}

func NewExpenseDao() ExpenseDao {
	return &expenseDao{
		expenseDao: make(map[int]entities.IExpense),
	}
}

func (e *expenseDao) CreateExpense(id int, description string, amount float64, createdBy entities.IUser) (entities.IExpense, error) {
	expense := entities.NewExpense(id, description, amount, createdBy)
	err := e.addExpense(expense)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (e *expenseDao) addExpense(expense entities.IExpense) error {
	if _, ok := e.expenseDao[expense.GetId()]; ok {
		return errors.New("expense already exists")
	}
	e.expenseDao[expense.GetId()] = expense
	return nil
}

func (e *expenseDao) GetExpense(id int) (entities.IExpense, error) {
	if expense, ok := e.expenseDao[id]; ok {
		return expense, nil
	}
	return nil, errors.New("expense not found")
}

func (e *expenseDao) DeleteExpense(id int) {
	if _, ok := e.expenseDao[id]; ok {
		delete(e.expenseDao, id)
		return
	}
}

func (e *expenseDao) UpdateExpense(id int, description string, amount float64, createdBy entities.IUser) (entities.IExpense, error) {
	if expense, ok := e.expenseDao[id]; ok {
		expense.SetDescription(description).SetAmount(amount).SetCreatedBy(createdBy)
		return expense, nil
	}
	return nil, errors.New("expense not found")
}

func (e *expenseDao) UpdateExpenseByParam(paramName, val interface{}, id int) (entities.IExpense, error) {
	if expense, ok := e.expenseDao[id]; ok {
		switch paramName {
		case "description":
			expense.SetDescription(val.(string))
		case "amount":
			expense.SetAmount(val.(float64))
		case "createdBy":
			expense.SetCreatedBy(val.(entities.IUser))
		}
		return expense, nil
	}
	return nil, errors.New("expense not found")
}
