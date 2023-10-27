package entities

type IAccount interface {
	GetAccountId() int
	SetWalletId(walletId int) IAccount
	GetBalance() float64
	SetBalance(balance float64) IAccount
	GetUserId() int
	SetUserId(userId int) IAccount
	AddTransaction(transaction ITransaction) IAccount
	GetTransactions() []ITransaction
}

type account struct {
	walletId     int
	balance      float64
	userId       int
	transactions []ITransaction
}

func NewDefaultAccount() IAccount {
	return &account{}
}

func NewAccount(walletId int, balance float64, userId int) IAccount {
	return &account{
		walletId:     walletId,
		balance:      balance,
		userId:       userId,
		transactions: make([]ITransaction, 0),
	}
}

func (w *account) GetAccountId() int {
	return w.walletId
}

func (w *account) SetWalletId(walletId int) IAccount {
	w.walletId = walletId
	return w
}

func (w *account) GetBalance() float64 {
	return w.balance
}

func (w *account) SetBalance(balance float64) IAccount {
	w.balance = balance
	return w
}

func (w *account) GetUserId() int {
	return w.userId
}

func (w *account) SetUserId(userId int) IAccount {
	w.userId = userId
	return w
}

func (w *account) GetTransactions() []ITransaction {
	return w.transactions
}

func (w *account) AddTransaction(transaction ITransaction) IAccount {
	w.transactions = append(w.transactions, transaction)
	return w
}
