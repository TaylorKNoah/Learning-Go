package account

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
