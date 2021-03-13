package cash

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Deposit $50 into Alice's account
func TestDeposit(t *testing.T) {
	account := account{
		"Alice",
		50,
	}
	account = deposit(account, 50)
	assert.Equal(t, 100, account.balance)
}

func TestWithdrawInsufficientBalance(t *testing.T) {
	account := account{
		"Maya",
		50,
	}

	account, error := withdraw(account,60)
	assert.Errorf(t, error,"%s")
	assert.NotEqual(t, account.balance, 10)
}

func TestWithdraw(t *testing.T) {
	account := account{
		"Daisy",
		50,
	}
	account, error := withdraw(account, 10)
	assert.Equal(t, 40, account.balance)
	assert.NotEqual(t, error, "%s")
}

func TestFindOrCreateAccount(t *testing.T){
	//test create account
	account, error := findOrCreateAccounts("Jane",0)
	if error != nil {
		log.Fatal(error)
	}
	assert.Equal(t,account.customer,"Jane")
	assert.Equal(t, account.balance, 0)
	//test find account
	account, error = findOrCreateAccounts("Jane",50)
	assert.Equal(t,account.customer,"Jane")
	assert.Equal(t, account.balance, 50)
}

func TestEndtoEndWorkflow(t *testing.T){
	account := account{
		"Marley",
		0,
	}
	account = deposit(account,50)
	assert.Equal(t, 50,account.balance)
	account, error := withdraw(account,40)
	assert.Equal(t, 10,account.balance )
	assert.NotEqual(t,error,"Actual %s" )
	account, error = withdraw(account,40)
	assert.Errorf(t,error,"Actual %s" )
}

func TestLedger(t *testing.T){
	bankBalance = printLedgerAndShowBalance()
	assert.Equal(t, bankBalance, 200)
}

//func TestDoubleDeposit(t *testing.T){
//	accounts = append(accounts,
//		account{
//			"Sarah",
//			0,
//		},
//	)
//	account, err := findOrCreateAccounts("Sarah")
//	if err != nil {
//		t.Fatal("expected no error got", err)
//	}
//	transaction := deposit(account, 50)
//	transaction := deposit(account, 50)
//	assert.Equal(t, 100,transaction.account.balance)
//}
//
//
//
//
////func TestSetup(t *testing.T) {
////	accounts = append(accounts,
////		account{
////			"Alice",
////			50,
////		},
////		account{
////			"Daisy",
////			0,
////		},
////		account{
////			"Rose",
////			0,
////		},
////	)
////	fmt.Println(accounts)
////}
