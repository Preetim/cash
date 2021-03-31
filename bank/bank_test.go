package bank

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestDepositWithBalance(t *testing.T) {
//	account := CreateNewAccount("Alice")
//	account = Deposit(account, 50)
//	account = Deposit(account, 50)
//	assert.Equal(t, 100, account.balance)
//}
//
//func TestNewDeposit(t *testing.T){
//	account := CreateNewAccount("Dasha")
//	account = Deposit(account,10)
//	assert.Equal(t, 10,CalculateAccountBalance(account))
//}
//
//func TestWithdrawInsufficientBalance(t *testing.T) {
//	account := CreateNewAccount("Dasha")
//	account = Deposit(account,50)
//	account, error := Withdraw(account,60)
//	assert.Equal(t, 50, CalculateAccountBalance(account))
//	assert.NotEqual(t, error,"Insufficient balance")
//}
//
//func TestWithdraw(t *testing.T) {
//	account := CreateNewAccount("Dasha")
//	account = Deposit(account,50)
//	account, error := Withdraw(account, 10)
//	assert.Equal(t, 40, CalculateAccountBalance(account))
//	assert.NotEqual(t, error, "%s")
//}
//
//func TestWithdrawFromZeroBalance(t *testing.T){
//	account := CreateNewAccount("Rini")
//	account, error := Withdraw(account, 5)
//	assert.Equal(t, 0, CalculateAccountBalance(account))
//	assert.Errorf(t, error,"%s")
//}
//
//func TestEndtoEndWorkflow(t *testing.T){
//	account := CreateNewAccount("Marley")
//	account = Deposit(account,50)
//	assert.Equal(t, 50,CalculateAccountBalance(account))
//	account, error := Withdraw(account,40)
//	assert.Equal(t, 10,CalculateAccountBalance(account))
//	assert.NotEqual(t,error,"Actual %s" )
//	account, error = Withdraw(account,40)
//	assert.Errorf(t,error,"Actual %s" )
//}
//
//func TestBankBalance(t *testing.T){
//	account := CreateNewAccount("Marley")
//	account = Deposit(account,50)
//	account2 := CreateNewAccount("Charlie")
//	account2 = Deposit(account2,70)
//	accounts := []Account{account,account2}
//	bankBalance := CalculateBankBalance(accounts)
//	assert.Equal(t, 120, bankBalance)
//}


func TestOverDraft(t *testing.T){
	var accounts =  []Account {}
	PreeAccount := CreateNewAccount("Preeti")
	PreeAccount, error := Withdraw(PreeAccount, accounts, 10)
	assert.Errorf(t, error, "%s")
	JimAccount := CreateNewAccount("Jim")
	JimAccount = Deposit(JimAccount, 50)
	accounts = []Account{PreeAccount,JimAccount}
	PreeAccount, error = Withdraw(PreeAccount,accounts ,10)
	assert.Equal(t, -10,PreeAccount.balance)

}