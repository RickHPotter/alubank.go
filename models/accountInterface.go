package models

/*
USER INTERFACE
*/

// DATA

func IGetName(account verify) string {
	return account.GetName()
}

func IGetAgNum(account verify) int {
	return account.GetAgNum()
}

func IGetAccNum(account verify) int {
	return account.GetAccNum()
}

func ICheckBalance(account verify) string {
	return account.CheckBalance()
}

// OPERATIONS

func IWithdraw(account verify, number float64) {
	account.Withdraw(number)
}

func IDeposit(account verify, number float64) {
	account.Deposit(number)
}

func PayBill(account verify, barcode uint16, number float64) {
	account.PayBill(barcode, number)
}

type verify interface {
	GetName() string
	GetAgNum() int
	GetAccNum() int
	CheckBalance() string

	Withdraw(number float64)
	Deposit(number float64)
	PayBill(barcode uint16, number float64)
}
