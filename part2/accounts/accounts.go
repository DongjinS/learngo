package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type account struct {
	owner string;
	balance int;
}

var errNoMoney = errors.New("can't withdraw you are poor")

// NewAccount creates account
func NewAccount(owner string) *account {
	account := account{owner: "DJ", balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *account) Deposit(amount int) {
	a.balance+=amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}

// withdraw x amount of your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance-=amount
	return nil
}

// ChangeOwner of your account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of your account
func (a account) Owner() string {
	return a.owner
}

func (a account) String() string {
	// Sprint make string
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}