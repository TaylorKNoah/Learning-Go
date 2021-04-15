package main

import (
	"fmt"

	"github.com/taylorknoah/Learning-Go/BankProject/account"
)

func main() {

	//This block creats an account, adds to the balance, and displays the balance
	my_account := account.NewAccount("taylor")
	my_account.Deposit(10)
	fmt.Println(my_account.Balance())

	//This block trys to withdraw from the balance
	//If the withdraw is greater than the balance an error is returned and displayed
	//Displays balance
	errWithdraw := my_account.Withdraw(20)
	if errWithdraw != nil {
		fmt.Println(errWithdraw)
	}
	fmt.Println(my_account.Balance())
}
