package p6_pointers_n_errors

import (
	"fmt"
	"testing"
)

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
			t.Fatal(fmt.Sprintf("got error %v when didnt expect one", err))
		}
	}

	assertError := func(t testing.TB, gotErr error, wantMsg string) {
		t.Helper()

		if gotErr == nil {
			t.Fatal("wanted an error but didnt get one")
		}
		if gotErr.Error() != wantMsg {
			t.Errorf("got %q, want %q", gotErr.Error(), wantMsg)
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

		withdrawalBalance := Bitcoin(100)
		err := wallet.Withdraw(withdrawalBalance)
		want := initialBalance

		assertBalance(t, wallet, want)
		assertError(t, err, fmt.Sprintf("Insufficient balance, available %s, withdrawing %s", initialBalance, withdrawalBalance))
	})
}
