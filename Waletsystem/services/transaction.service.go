package services

import (
	"errors"
	"pattern/Waletsystem/dao"
	"pattern/Waletsystem/entities"
)

type ITransactionService interface {
	TransferMoney(amount float64, account1, account2 entities.IAccount) error
	CreditMoney(account entities.IAccount, amount float64) error
	DebitMoney(account entities.IAccount, amount float64) error
	GetTransactionByTransactionId(transactionId int) error
}

func NewTransactionService() ITransactionService {
	return &transactionService{
		transactionDao: dao.NewTransactionDAO(),
	}
}

type transactionService struct {
	transactionDao dao.ITransactionDAO
}

func (t *transactionService) TransferMoney(amount float64, fromAccount, toAccount entities.IAccount) error {

	transaction1 := entities.NewTransaction(1, amount, fromAccount.GetAccountId(), entities.DEBIT)
	transaction2 := entities.NewTransaction(2, amount, fromAccount.GetAccountId(), entities.CREDIT)

	if fromAccount.GetBalance() >= amount {
		fromAccount.SetBalance(fromAccount.GetBalance() - amount).AddTransaction(transaction1)
		toAccount.SetBalance(toAccount.GetBalance() + amount).AddTransaction(transaction2)

		transaction1.SetTransactionStatus(entities.COMPLETED)
		transaction2.SetTransactionStatus(entities.COMPLETED)

		t.transactionDao.InsertTransactionByTransactionId(transaction1)
		t.transactionDao.InsertTransactionByTransactionId(transaction2)
		return nil
	}
	transaction1.SetTransactionStatus(entities.CANCELLED)
	fromAccount.AddTransaction(transaction1)
	t.transactionDao.InsertTransactionByTransactionId(transaction1)
	return errors.New("insufficient Balance")
}

func (t *transactionService) CreditMoney(account entities.IAccount, amount float64) error {
	transaction := entities.NewTransaction(1, amount, account.GetAccountId(), entities.CREDIT)
	account.SetBalance(account.GetBalance() + amount)
	transaction.SetTransactionStatus(entities.COMPLETED)
	account.AddTransaction(transaction)
	t.transactionDao.InsertTransactionByTransactionId(transaction)
	return nil
}

func (t *transactionService) DebitMoney(account entities.IAccount, amount float64) error {
	transaction := entities.NewTransaction(1, amount, account.GetAccountId(), entities.DEBIT)
	if account.GetBalance() >= amount {
		account.SetBalance(account.GetBalance() - amount)
		transaction.SetTransactionStatus(entities.COMPLETED)
		account.AddTransaction(transaction)
		t.transactionDao.InsertTransactionByTransactionId(transaction)
		return nil
	}
	transaction.SetTransactionStatus(entities.CANCELLED)
	account.AddTransaction(transaction)
	t.transactionDao.InsertTransactionByTransactionId(transaction)
	return errors.New("insufficient Balance")
}

func (t *transactionService) GetTransactionByTransactionId(transactionId int) error {
	return t.transactionDao.GetTransactionByTransactionId(transactionId)
}