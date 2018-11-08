package pointers

import "fmt"

type Bitcoin int

type WalletErr string

const ErrInsufficientFunds = WalletErr("cannot withdraw, insuficient funds")

func (e WalletErr) Error() string {
	return string(e)
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin	
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

