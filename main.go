package main

import (
	"fmt"
	"github/RickHPotter/alubank_alura_course/models"
)

func main() {

	acc := models.NewAccount("George", 123, 456, 2324.66)
	// acc2 := models.Account{Name: "George"}
	// acc3 := models.Account{Name: "George"}

	fmt.Println(*acc)
	models.Withdraw(acc, 324.66)
	fmt.Println(*acc)
	models.Withdraw(acc, 1000)
	fmt.Println(*acc)
	models.Withdraw(acc, 980)
	fmt.Println(*acc)
}
