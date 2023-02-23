package main

import (
	"fmt"
	"github/RickHPotter/alubank_alura_course/models"
)

func main() {

	source_acc := models.NewAccount("Perl", 123, 456, 2424.66)
	target_acc := models.NewAccount("Julia", 123, 456, 2424.66)

	fmt.Println(source_acc.GetName(), source_acc.CheckBalance())
	fmt.Println(target_acc.GetName(), target_acc.CheckBalance())
	source_acc.Transfer(target_acc, 0.66)
	fmt.Println(source_acc.GetName(), source_acc.CheckBalance())
	fmt.Println(target_acc.GetName(), target_acc.CheckBalance())
	fmt.Println()
	target_acc.Withdraw(2)
	fmt.Println(source_acc.GetName(), source_acc.CheckBalance())
	fmt.Println(target_acc.GetName(), target_acc.CheckBalance())

}
