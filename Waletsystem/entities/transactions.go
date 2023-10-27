package entities

import "time"

type ITransaction interface {
	GetTransactionId() int
	SetTransactionId(transactionId int) ITransaction
	GetAmount() float64
	SetAmount(amount float64) ITransaction
	GetFromWalletId() int
	SetFromWalletId(fromWalletId int) ITransaction
	GetTranasctionType() TransactionType
	SetTranasctionType(tranadctionType TransactionType) ITransaction
	GetTransactionTime() time.Time
	SetTransactionTime(transactionTime time.Time) ITransaction
	GetTransactionStatus() TransactionStatus
	SetTransactionStatus(transactionStatus TransactionStatus) ITransaction
}

type transaction struct {
	transactionId     int
	amount            float64
	fromWalletId      int
	transactionType   TransactionType
	transactionTime   time.Time
	transactionStatus TransactionStatus
}

func NewTransaction(transactionId int, amount float64, fromWalletId int, transactionType TransactionType) ITransaction {
	return &transaction{
		transactionId:   transactionId,
		amount:          amount,
		fromWalletId:    fromWalletId,
		transactionType: transactionType,
		transactionTime: time.Now(),
	}
}

func (t *transaction) GetTransactionId() int {
	return t.transactionId
}

func (t *transaction) SetTransactionId(transactionId int) ITransaction {
	t.transactionId = transactionId
	return t
}

func (t *transaction) GetAmount() float64 {
	return t.amount
}

func (t *transaction) SetAmount(amount float64) ITransaction {
	t.amount = amount
	return t
}

func (t *transaction) GetFromWalletId() int {
	return t.fromWalletId
}

func (t *transaction) SetFromWalletId(fromWalletId int) ITransaction {
	t.fromWalletId = fromWalletId
	return t
}

func (t *transaction) GetTranasctionType() TransactionType {
	return t.transactionType
}

func (t *transaction) SetTranasctionType(tranadctionType TransactionType) ITransaction {
	t.transactionType = tranadctionType
	return t
}

func (t *transaction) GetTransactionTime() time.Time {
	return t.transactionTime
}

func (t *transaction) SetTransactionTime(transactionTime time.Time) ITransaction {
	t.transactionTime = transactionTime
	return t
}

func (t *transaction) GetTransactionStatus() TransactionStatus {
	return t.transactionStatus
}

func (t *transaction) SetTransactionStatus(transactionStatus TransactionStatus) ITransaction {
	t.transactionStatus = transactionStatus
	return t
}
