package main

import "fmt"
import "errors"

type Bitcoin float64

type Stringer interface {
	String() string
}

// Wallet stores the bitcoin Wallet state, such as a/c balance
type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("Insufficient funds, cannot withdraw")

// Deposit adds the amount to the wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the current Wallet balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw deducts amount from account
func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
