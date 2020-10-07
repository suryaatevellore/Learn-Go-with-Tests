package main

import "testing"

func TestWallet(t *testing.T) {
	// t.Fatal("not implemented")
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got.String(), want.String())
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but didn't get one")
		}
		if got.Error() != want.Error() {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("Did not want an error but got one")
		}
	}

	t.Run("test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("test Withdrawal", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)
		err := wallet.Withdraw(10)
		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("test Withdrawal from an empty account", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		want := ErrInsufficientFunds
		assertError(t, err, want)
	})
}
