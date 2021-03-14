package cash

import (
	"github.com/stretchr/testify/assert"
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

//Deposit money into a new customer's account
func TestNewDeposit(t *testing.T){
	//var a account
	account := account{
		"Alice",
		0,
	}
	account = deposit(account,10)
	assert.Equal(t, 10,account.balance)
}

//Withdraw insufficient balance 60 from 50 balance
func TestWithdrawInsufficientBalance(t *testing.T) {
	account := account{
		"Maya",
		50,
	}
	account, error := withdraw(account,60)
	assert.Errorf(t, error,"%s")
	assert.NotEqual(t, account.balance, -10)
}

//Test withdraw 10 from 50
func TestWithdraw(t *testing.T) {
	account := account{
		"Daisy",
		50,
	}
	account, error := withdraw(account, 10)
	assert.Equal(t, 40, account.balance)
	assert.NotEqual(t, error, "%s")
}

//Test withdraw from 0 balance
func TestWithdrawFromZeroBalance(t *testing.T){
	var a account
	a.customer = "Rini"
	a, error := withdraw(a, 5)
	assert.Equal(t, 0, a.balance)
	assert.Errorf(t, error,"%s")
}

//Test deposit, withdraw and then withdraw with insufficient funds
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

func TestBankBalance(t *testing.T){
	accounts = [] account{
		{
			"Daisy",
			50,
		},
		{
			"Mor",
			70,
		},
	}
	bankBalance = calculateBankBalance()
	assert.Equal(t, 120,bankBalance)
}

func TestFindAccount(t *testing.T){
	accounts = [] account{
		{
			"Daisy",
			50,
		},
		{
			"Mor",
			70,
		},
	}
	account := findAccount("Mor")
	assert.Equal(t, "Mor",account.customer)
}

//Test end of day settlement files and bank balance. In a real world this is an output every day for banks
func TestPrintSettlementFile(t *testing.T){
	printLedger()
}