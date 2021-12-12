package p6_pointers_n_errors

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(btc Bitcoin) {
	w.balance += btc
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
