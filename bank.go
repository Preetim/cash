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
	bankBalance += amount

	return t
}


















