package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

//Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance return the number of Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw will remove some Bitcoin to a wallet.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
