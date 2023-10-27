package dao

import (
	"errors"
	"pattern/Waletsystem/entities"
)

type ITransactionDAO interface {
	InsertTransactionByTransactionId(transaction entities.ITransaction) error
	GetTransactionByTransactionId(transactionId int) error
}

func NewTransactionDAO() ITransactionDAO {
	return &transactionDAO{
		transactionIdDetailsMap: make(map[int]entities.ITransaction),
	}
}

type transactionDAO struct {
	transactionIdDetailsMap map[int]entities.ITransaction
}

func (t *transactionDAO) InsertTransactionByTransactionId(transaction entities.ITransaction) error {
	if _, ok := t.transactionIdDetailsMap[transaction.GetTransactionId()]; ok {
		return errors.New("transaction Id already exists")
	}
	t.transactionIdDetailsMap[transaction.GetTransactionId()] = transaction
	return nil
}

func (t *transactionDAO) GetTransactionByTransactionId(transactionId int) error {
	if _, ok := t.transactionIdDetailsMap[transactionId]; ok {
		return nil
	}
	return errors.New("transaction not found")
}
