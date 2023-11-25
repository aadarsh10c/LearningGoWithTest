package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()
		want := Bitcoin(30)

		checkBitcoin(wallet, want, t)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()
		want := Bitcoin(10)

		checkBitcoin(wallet, want, t)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		checkBitcoin(wallet, startingBalance, t)
		if err == nil {
			t.Errorf("Wanted an error bit got nil")
		}

	})

}

func checkBitcoin(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
