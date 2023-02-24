package models

import (
	"fmt"
	"github/RickHPotter/alubank_alura_course/utils"
	"strconv"
)

type Account struct {
	holder        Client
	agencyNumber  int
	accountNumber int
	balance       float64
}

/*
DEV INTERFACE
*/

func (self_acc *Account) _add(number float64) {
	res := utils.RoundFloat(self_acc.balance+number, 2)
	self_acc.balance = res
}

func (self_acc *Account) _subtract(number float64) {
	res := utils.RoundFloat(self_acc.balance-number, 2)
	self_acc.balance = res
}

func _calculateFees(operation string, number float64) float64 {
	var fee float64

	switch operation {
	case "withdraw":
		fee = utils.RoundFloat(number*0.01, 2)
	case "checkBalance":
		fee = 0.36
	}

	return fee
}

/*
USER INTERFACE
*/

// DATA

func (self_acc *Account) GetName() string {
	return self_acc.holder.name
}

func (self_acc *Account) GetAgNum() int {
	return self_acc.agencyNumber
}

func (self_acc *Account) GetAccNum() int {
	return self_acc.accountNumber
}

func (self_acc *Account) CheckBalance() string {
	balance := strconv.FormatFloat(self_acc.balance, 'f', 2, 64)
	return "£" + balance
	// TODO: implement fee for production
}

// OPERATIONS

func NewAccount(holder *Client, agencyNumber int, accountNumber int, balance float64) *Account {
	account := Account{*holder, agencyNumber, accountNumber, balance}
	return &account
}

func (self_acc *Account) Withdraw(withdraw float64) {
	fee := _calculateFees("withdraw", withdraw)

	if withdraw > 0 {

		if self_acc.balance >= withdraw+fee {
			self_acc._subtract(withdraw + fee)

			fmt.Print("[ACCEPTED]\n")
			fmt.Print("£", withdraw, " checked out of your account.\n")
			fmt.Print("£", fee, " was the operation fee.\n")
			fmt.Print("£", self_acc.balance, " is your new balance.\n")
		} else {
			if self_acc.balance >= withdraw {
				fmt.Println("[DENIED] Enough money to withdraw but not enough for fee.")
			} else {
				fmt.Println("[DENIED] You cannot withdraw more money than your balance.")
			}
		}
	} else {
		fmt.Println("[DENIED] (Negative Withdraw Not Allowed) Man, you must be tripping.")
	}
}

func (self_acc *Account) Deposit(deposit float64) {
	if deposit > 0 {
		self_acc._add(deposit)
		fmt.Println("[ACCEPTED] Done. Your new balance is", self_acc.balance)
	} else {
		fmt.Println("[DENIED] (Negative Deposit Not Allowed) Man, you must be tripping.")
	}
}

func (self_acc *Account) Transfer(target_acc *Account, number float64) {
	// let's consider that there's only one bank, therefore no fee for transfer
	if number <= self_acc.balance {
		self_acc._subtract(number)
		fmt.Print("£", number, " was reduced from ", self_acc.GetName(), "'s account,\n")
		target_acc._add(number)
		fmt.Print("£", number, " was added to ", target_acc.GetName(), "'s account.\n")
		fmt.Println("Handshake!")
	} else {
		fmt.Println("[DENIED] (Insufficient Funds) Man, you must be tripping.")
	}

}
