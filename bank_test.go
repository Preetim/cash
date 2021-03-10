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
			0,
		},
		account{
			"Rose",
			0,
		},
	)
	fmt.Println(accounts)
}

// Deposit $50 into Alice's account
func TestDeposit(t *testing.T) {
	account, err := findAccount("Alice")
	if err != nil {
		t.Fatal("expected no error got", err)
	}
	transaction := deposit(account, 50)
	if transaction.account.balance !=  100 {
		t.Errorf("got #{transaction.account.balance}; want 100")
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

func TestBankBalance(t *testing.T) {
	bankBalance := 0
	bankBalance = findBankBalance(transactions)
	if bankBalance !=  50 {
		t.Errorf("got #{bankBalance}; want 50")
	}
}







