package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// method to deposit money in wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance = w.balance + amount
}

// method to check balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var errInsuffucientfunds = errors.New("cannot withdraw, insufficient funds")

// method to withdraw
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsuffucientfunds
	}
	w.balance = w.balance - amount
	return nil
}
