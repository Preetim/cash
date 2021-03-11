package cash

import (
	"fmt"
)

type transaction struct {
	transactionType string
	amount int
	account account
}

type account struct {
	customer string
	balance int
}

var accounts []account
var transactions []transaction
var bankBalance int = 0

func deposit(account account, amount int) transaction {
	account.balance += amount
	t := transaction{
		transactionType: "deposit",
		amount: amount,
		account: account,
	}
	transactions = append(transactions, t)
	return t
}

func createAccount(customer string) account{
	newAccount := account{
		customer: customer,
		balance: 0,
	}
	accounts = append(accounts,newAccount)
	return newAccount
}

func findBankBalance(accounts []account) int {
	bankBalance = 0
	for i := range accounts {
		bankBalance = accounts[i].balance + bankBalance
		fmt.Println(accounts[i])
	}
	return bankBalance
}


func withdraw(account account, amount int) (transaction, error) {
	account.balance -= amount
	t := transaction{
		transactionType: "withdraw",
		amount: amount,
		account: account,
	}
	if account.balance < 0{
		return t, fmt.Errorf("Insufficient funds")
	}
	transactions = append(transactions, t)
	return t, nil
}


















