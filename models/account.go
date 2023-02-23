package models

import (
	"fmt"
	"github/RickHPotter/alubank_alura_course/utils"
)

type Account struct {
	Name          string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func NewAccount(name string, agencyNumber int, accountNumber int, balance float64) *Account {
	account := Account{name, agencyNumber, accountNumber, balance}
	return &account
}

func Withdraw(acc *Account, withdraw float64) {
	fee := utils.RoundFloat(_calculateFees("withdraw", withdraw), 2)
	newBalance := utils.RoundFloat(acc.Balance-withdraw-fee, 2)

	if newBalance >= 0 {
		acc.Balance = newBalance

		fmt.Print("[ACCEPTED] £", withdraw, " has been checked out of your account.\n")
		fmt.Print("The fee for this operation was £", fee, ".\n")
		fmt.Print("Your new balance is £", acc.Balance, ".\n")
	} else {
		if acc.Balance >= withdraw {
			fmt.Println("[DENIED] Enough money to withdraw but not enough for fee.")
		} else {
			fmt.Println("[DENIED] You cannot withdraw more money than your balance.")
		}
	}
}

func Deposit() {

}

func CheckBalance() {

}

func _calculateFees(operation string, number float64) float64 {
	var fee float64

	switch operation {
	case "withdraw":
		fee = number * 0.01
	case "deposit":
		fee = 0
	case "checkBalance":
		fee = 0.36
	}

	return fee
}
