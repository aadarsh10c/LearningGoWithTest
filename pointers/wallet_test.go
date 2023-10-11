package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{Bitcoin(20)}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(30)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
