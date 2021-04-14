package main

import (
	"fmt"

	"github.com/taylorknoah/Learning-Go/BankProject/Banking"
)

func main() {
	account := Banking.Account{Owner: "Taylor", Balance: 1000}
	fmt.Println(account)
}
