package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertEqual(got, want, t)
	})

	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(120))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertEqual(got, want, t)
		assertError(err, errInsuffucientfunds, t)

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(Bitcoin(120))
		got := wallet.Balance()
		want := Bitcoin(20)

		assertError(err, errInsuffucientfunds, t)
		assertEqual(got, want, t)

	})

}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but did n't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertEqual(got Bitcoin, want Bitcoin, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
