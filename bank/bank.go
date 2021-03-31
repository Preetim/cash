package bank

import "fmt"

type Account struct{
	customer string
	balance int
}

//var accounts []Account

func Deposit(account Account, amount int) Account {
	account.balance += amount
	return account
}

func Withdraw(account Account,accounts []Account, amount int) (Account, error) {
	bankBalance := CalculateBankBalance(accounts)
	if bankBalance < amount{
		return account, fmt.Errorf("Insufficient funds")
	} else {
		account.balance -= amount
	}
	return account,nil
}

func CalculateAccountBalance(account Account) int {
	return account.balance
}

func CreateNewAccount(customerName string) Account{
	a := Account{customer: customerName, balance: 0,}
	//accounts = append(accounts,a)
	return a
}

func CalculateBankBalance(bank []Account) int {
	bankBalance := 0
	for _, i := range bank{
		bankBalance += i.balance
	}
	return bankBalance
}

//func FindAccount(customerName string) (Account, error){
//	for _, i := range accounts{
//		if i.customer == customerName{
//			return i, nil
//		}
//	}
//	return Account{}, fmt.Errorf("Account not found")
//}
