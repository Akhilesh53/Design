package services

import (
	"pattern/Waletsystem/dao"
	"pattern/Waletsystem/entities"
)

type IAccountSerivce interface {
	CreateAccount(user entities.IUser) entities.IAccount
	GetAccountByAccountId(accountId int) (entities.IAccount, error)
	UpdateAccountByAccountId(accountId int, account entities.IAccount) IAccountSerivce
}

type accountService struct {
	accountDao         dao.IAccountDAO
	accountIdGenerator entities.IAccountNumberGenerator
}

func NewAccountService() IAccountSerivce {
	return &accountService{
		accountDao:         dao.NewWalletDAO(),
		accountIdGenerator: entities.NewAccountNumberGenerator(),
	}
}

func (a *accountService) CreateAccount(user entities.IUser) entities.IAccount {
	account := entities.NewAccount(a.accountIdGenerator.Generate(), 0, user.GetUserId())
	a.accountDao.InsertAccountByAccountId(account.GetAccountId(), account)
	return account
}

func (a *accountService) GetAccountByAccountId(accountId int) (entities.IAccount, error) {
	return a.accountDao.FindWalletByWalletId(accountId)
}

func (a *accountService) UpdateAccountByAccountId(accountId int, account entities.IAccount) IAccountSerivce {
	a.accountDao.UpdateWalletByWalletId(accountId, account)
	return a
}
