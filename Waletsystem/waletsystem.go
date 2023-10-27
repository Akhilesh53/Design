package main

import (
	"fmt"
	"pattern/Waletsystem/services"
)

func main() {

	userService := services.NewUserService()
	transactionService := services.NewTransactionService()
	accountService := services.NewAccountService()

	user1 := userService.CreateUser("John")
	fmt.Println(user1)
	account1 := accountService.CreateAccount(user1)
	fmt.Println(account1)

	user2 := userService.CreateUser("Rahul")
	fmt.Println(user2)
	account2 := accountService.CreateAccount(user2)
	fmt.Println(account2)
	
	account1.SetBalance(1000)

	if err := transactionService.TransferMoney(100, account1, account2); err != nil {
		fmt.Println(err)
	}

	for _, transaction := range account1.GetTransactions() {
		fmt.Println(transaction.GetTransactionId(), transaction.GetTransactionStatus())
	}

	for _, transaction := range account2.GetTransactions() {
		fmt.Println(transaction.GetTransactionId(), transaction.GetTransactionStatus())
	}

	fmt.Println(account1.GetBalance())
	fmt.Println(account2.GetBalance())
}
