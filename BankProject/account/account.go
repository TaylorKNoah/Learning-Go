package account

import "errors"

//Account Struct
type account struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *account {
	new_Account := account{owner: owner, balance: 0}
	return &new_Account
}

// Deposit x amount to the account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Return balance
func (a account) Balance() int {
	return a.balance
}

//withdraw from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw. you are poor")
	}
	a.balance -= amount
	return nil
}
