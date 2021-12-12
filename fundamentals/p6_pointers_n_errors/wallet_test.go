package p6_pointers_n_errors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("in %#v, got %s, want %s", wallet, got, want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("got error %v when didnt expect one", err)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err == nil {
			t.Errorf("wanted an error but didnt get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw: normal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(5)
		want := Bitcoin(15)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw: insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}

		err := wallet.Withdraw(Bitcoin(100))
		want := initialBalance

		assertBalance(t, wallet, want)
		assertError(t, err)
	})
}
