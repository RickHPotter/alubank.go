package main

import (
	"fmt"
	"github/RickHPotter/alubank_alura_course/models"
)

func main() {

	p1 := models.NewClient("Ruby", "111", "Engineer")
	p2 := models.NewClient("Julia", "222", "Scientist")

	source_acc := models.NewBankAccount(p1, 123, 456, 1, 2424.66)
	target_acc := models.NewBankAccount(p2, 123, 456, 1, 2424.66)

	models.IGetName(source_acc)

	fmt.Println(models.IGetName(source_acc), source_acc.CheckBalance())
	fmt.Println(models.IGetName(target_acc), target_acc.CheckBalance())
	// source_acc.Transfer(target_acc, 0.66)
	// fmt.Println(source_acc.GetName(), source_acc.CheckBalance())
	// fmt.Println(target_acc.GetName(), target_acc.CheckBalance())
	// fmt.Println()
	// target_acc.Withdraw(2)
	// fmt.Println(source_acc.GetName(), source_acc.CheckBalance())
	// fmt.Println(target_acc.GetName(), target_acc.CheckBalance())

}
