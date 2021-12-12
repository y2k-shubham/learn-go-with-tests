package p6_pointers_n_errors

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(btc int) {
	w.balance += btc
}

func (w *Wallet) Balance() int {
	return w.balance
}
