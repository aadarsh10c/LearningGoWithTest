package wallet

type Wallet struct {
	balance int
}

// method to deposit money in wallet
func (w *Wallet) Deposit(amount int) {
	w.balance = w.balance + amount
}

// method to check balance
func (w *Wallet) Balance() int {
	return w.balance
}
