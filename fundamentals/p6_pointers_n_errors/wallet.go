package p6_pointers_n_errors

import "fmt"

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

func (w *Wallet) Withdraw(btc Bitcoin) {
	w.balance -= btc
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
