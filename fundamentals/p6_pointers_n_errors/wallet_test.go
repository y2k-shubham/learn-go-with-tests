package p6_pointers_n_errors

import "testing"

func TestWallet(t *testing.T) {
	checkBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("in %#v, got %s, want %s", wallet, got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)

		checkBalance(t, wallet, want)
	})

	t.Run("Withdraw: normal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(5)
		want := Bitcoin(15)

		checkBalance(t, wallet, want)
	})

	t.Run("Withdraw: insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		err := wallet.Withdraw(Bitcoin(100))
		want := initialBalance

		checkBalance(t, wallet, want)

		if err == nil {
			t.Errorf("wanted an error but didnt get one")
		}
	})
}
