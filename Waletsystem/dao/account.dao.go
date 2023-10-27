package dao

import (
	"errors"
	"pattern/Waletsystem/entities"
)

type IAccountDAO interface {
	FindWalletByWalletId(walletId int) (entities.IAccount, error)
	InsertAccountByAccountId(walletId int, wallet entities.IAccount)
	UpdateWalletByWalletId(walletId int, wallet entities.IAccount) IAccountDAO
}

type accountDAO struct {
	walletIdDetailsMap map[int]entities.IAccount
}

func NewWalletDAO() IAccountDAO {
	return &accountDAO{
		walletIdDetailsMap: make(map[int]entities.IAccount),
	}
}

func (w *accountDAO) FindWalletByWalletId(walletId int) (entities.IAccount, error) {
	if wallet, ok := w.walletIdDetailsMap[walletId]; ok {
		return wallet, nil
	}
	return nil, errors.New("wallet id not found")
}

func (w *accountDAO) InsertAccountByAccountId(walletId int, wallet entities.IAccount)  {
	w.walletIdDetailsMap[walletId] = wallet
}

func (w *accountDAO) UpdateWalletByWalletId(walletId int, wallet entities.IAccount) IAccountDAO {
	w.walletIdDetailsMap[walletId] = wallet
	return w
}
