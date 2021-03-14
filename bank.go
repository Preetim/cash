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

func deposit(account account, amount int) account {
	account.balance += amount
	t := transaction{
		transactionType: "deposit",
		amount: amount,
		account: account,
	}
	transactions = append(transactions, t)
	return account
}

func withdraw(account account, amount int) (account, error) {
	if account.balance < amount{
		return account, fmt.Errorf("Insufficient funds")
	}
	account.balance -= amount
	t := transaction{
		transactionType: "withdraw",
		amount: amount,
		account: account,
	}
	transactions = append(transactions, t)
	return account, nil
}

func findAccount(customerName string) account {
	for i := range accounts {
			if accounts[i].customer == customerName {
			return accounts[i]
		}
	}
	return account{"",0}
}

func calculateBankBalance() int {
	for i := range accounts {
		fmt.Println(accounts[i])
		bankBalance += accounts[i].balance
	}
	return bankBalance
}

func printLedger() {
	for i := range transactions {
		fmt.Println(transactions[i])
	}
}