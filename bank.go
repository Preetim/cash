package cash

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

//func findBankBalance(transactions []transaction) int {
//	bankBalance = 0
//	for k, v := range transactions.account {
//		//fmt.Printf("Key is %d\n", k)
//		bankBalance += v.amount
//	}
//	fmt.Printf("bankBalance is %d\n", bankBalance)
//	return bankBalance
//}

func withdraw(account account, amount int) transaction {
	account.balance -= amount
	t := transaction{
		transactionType: "withdraw",
		amount: amount,
		account: account,
	}
	transactions = append(transactions, t)
	return t
}


















