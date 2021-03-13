package cash

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// Deposit $50 into Alice's account which has 50 balance
func TestDepositWithBalance(t *testing.T) {
	account := account{
		"Alice",
		50,
	}
	account = deposit(account, 50)
	assert.Equal(t, 100, account.balance)
}

func TestNewDeposit(t *testing.T){
	var a account
	a.customer = "Omaya"
	a = deposit(a,10)
	assert.Equal(t, 10,a.balance)
}

func TestWithdrawInsufficientBalance(t *testing.T) {
	account := account{
		"Maya",
		50,
	}
	account, error := withdraw(account,60)
	assert.Errorf(t, error,"%s")
	assert.NotEqual(t, account.balance, -10)
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

func TestWithdrawFromZeroBalance(t *testing.T){
	var a account
	a, error := withdraw(a, 5)
	assert.Equal(t, 0, a.balance)
	assert.Errorf(t, error,"%s")
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

func TestLedgerAndPrintSettltmentFile(t *testing.T){
	bankBalance = printLedgerAndShowBalance()
	assert.Equal(t, bankBalance, 210)
}
