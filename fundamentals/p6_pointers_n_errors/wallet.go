package p6_pointers_n_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(btc Bitcoin) {
	w.balance += btc
}

func (w *Wallet) Withdraw(btc Bitcoin) error {
	if w.balance < btc {
		return errors.New(fmt.Sprintf("Insufficient balance, available %s, withdrawing %s", w.Balance(), btc))
	}

	w.balance -= btc
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
