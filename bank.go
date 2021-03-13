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
	account, _ = findOrCreateAccounts(account.customer, account.balance)
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
	account, _ = findOrCreateAccounts(account.customer, account.balance)
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

func findOrCreateAccounts(customerName string, balance int) (account, error) {
	a :=account{
		customer: customerName,
		balance: balance,
	}
	for i := range accounts {
		if accounts[i].customer == customerName {
			accounts[i].balance = balance
			a = accounts[i]
			//return accounts[i], nil
		} else {
			newAccount := account{
				customer: customerName,
				balance:  0,
			}
			a = newAccount
			accounts = append(accounts, a)
		}
	}
	return a, nil
}

func printLedgerAndShowBalance() int {
	bankBalance := 0
	for i := range transactions {
		bankBalance += transactions[i].account.balance
		fmt.Println(transactions[i])
	}
	fmt.Printf("Total bank balance EOD is %d", bankBalance)
	return bankBalance
}
