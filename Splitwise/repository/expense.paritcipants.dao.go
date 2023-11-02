package repository

import "errors"

type ExpenseParticipantsDao interface {
}

type expenseParticipantsDao struct {
	expenseParticipantsDao map[int][]int // map<expenseId, [] userId>
}

func NewExpenseParticipantsDao() ExpenseParticipantsDao {
	return &expenseParticipantsDao{
		expenseParticipantsDao: make(map[int][]int),
	}
}

func (e *expenseParticipantsDao) AddExpenseParticipant(expenseId, userId int) error {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		e.expenseParticipantsDao[expenseId] = append(e.expenseParticipantsDao[expenseId], userId)
		return nil
	}
	return errors.New("expense not found")
}

func (e *expenseParticipantsDao) GetExpenseParticipants(expenseId int) ([]int, error) {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		return e.expenseParticipantsDao[expenseId], nil
	}
	return nil, errors.New("expense not found")
}

func (e *expenseParticipantsDao) GetExpenseParticipant(expenseId, userId int) (int, error) {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		for _, id := range e.expenseParticipantsDao[expenseId] {
			if id == userId {
				return id, nil
			}
		}
		return 0, errors.New("user not found")
	}
	return 0, errors.New("expense not found")
}

func (e *expenseParticipantsDao) UpdateExpenseParticipants(expenseId int, userIds []int) error {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		e.expenseParticipantsDao[expenseId] = userIds
		return nil
	}
	return errors.New("expense not found")
}

func (e *expenseParticipantsDao) DeleteExpense(expenseId int) {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		delete(e.expenseParticipantsDao, expenseId)
		return
	}
}

func (e *expenseParticipantsDao) DeleteExpenseParticipant(expenseId, userId int) {
	if _, ok := e.expenseParticipantsDao[expenseId]; ok {
		for i, id := range e.expenseParticipantsDao[expenseId] {
			if id == userId {
				e.expenseParticipantsDao[expenseId] = append(e.expenseParticipantsDao[expenseId][:i], e.expenseParticipantsDao[expenseId][i+1:]...)
				return
			}
		}
		return
	}
}
