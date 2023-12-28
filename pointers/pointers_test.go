package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(wallet Wallet, want Bitcoin, t testing.TB) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(wallet, Bitcoin(10), t)
	})
}
