package main

import (
	"fmt"

	"github.com/taylorknoah/Learning-Go/BankProject/account"
)

func main() {
	my_account := account.NewAccount("taylor")
	fmt.Println(my_account)
}
