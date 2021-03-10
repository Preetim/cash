package cash

import (
	"fmt"
	"testing"
)



func TestSetup(t *testing.T) {
	accounts = append(accounts,
		account{
			"Alice",
			50,
		},
		account{
			"Daisy",
			20,
		},
		account{
			"Rose",
			30,
		},
	)
	fmt.Println(accounts)
	bankBalance = 100 //Total of account balances
}

func TestDeposit(t *testing.T) {
	account, err := findAccount("Alice")
	if err != nil {
		t.Fatal("expected no error got", err)
	}
	transaction := deposit(account, 50)
	if transaction.account.balance !=  100 {
		t.Errorf("got #{transaction.account.balance}; want 100")
	}
	if bankBalance !=  150 {
		t.Errorf("Got #{bankBalance}; want 150")
	}
}








func findAccount(customerName string) (account, error) {
	for i := range accounts {
		//fmt.Printf("Accounts is %s", accounts[i].customer)
		if accounts[i].customer == customerName {
			return accounts[i],nil
		}

	}
	return account{"",0},fmt.Errorf("Account not found")
}







