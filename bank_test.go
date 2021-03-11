package cash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Deposit $50 into Alice's account
func TestDeposit(t *testing.T) {
	//setup data
	accounts = append(accounts,
		account{
			"Alice",
			50,
		},
	)
	account, err := findAccount("Alice")
	if err != nil {
		t.Fatal("expected no error got", err)
	}
	transaction := deposit(account, 50)
	assert.Equal(t, transaction.account.balance,100)
}

func TestWithdraw(t *testing.T) {
	//setup data
	accounts = append(accounts,
		account{
			"Alice",
			50,
		},
	)
	account, err := findAccount("Alice")
	if err != nil {
		t.Fatal("expected no error got", err)
	}
	transaction, error := withdraw(account,40)
	assert.Equal(t, transaction.account.balance,10)
	assert.NotEqual(t, error,"%s")
}

func TestWithdrawInsufficientBalance(t *testing.T) {
	accounts = append(accounts,
		account{
			"Alice",
			50,
		},
	)
	account, err := findAccount("Alice")
	if err != nil {
		t.Fatal("expected no error got", err)
	}
	transaction, error := withdraw(account,60)
	assert.Errorf(t, error,"%s")
	assert.NotEqual(t, transaction.account.balance, 10)
}

func findAccount(customerName string) (account, error) {
	for i := range accounts {
		if accounts[i].customer == customerName {
			return accounts[i],nil
		}
	}
	return account{"",0},fmt.Errorf("Account not found")
}

//func TestBankBalance(t *testing.T) {
//	accounts = append(accounts,
//		account{
//			"Alice",
//			50,
//			},
//		account{
//			"Daisy",
//			50,
//			},
//		)
//	bankBalance := 0
//	bankBalance = findBankBalance(transactions)
//	if bankBalance !=  100 {
//		t.Errorf("got #{bankBalance}; want 100")
//	}
//}


//func TestSetup(t *testing.T) {
//	accounts = append(accounts,
//		account{
//			"Alice",
//			50,
//		},
//		account{
//			"Daisy",
//			0,
//		},
//		account{
//			"Rose",
//			0,
//		},
//	)
//	fmt.Println(accounts)
//}









