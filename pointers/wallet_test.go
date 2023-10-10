package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{20}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 30

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
