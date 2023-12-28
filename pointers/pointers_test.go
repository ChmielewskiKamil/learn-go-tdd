package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

        assertError(err, ErrInsufficientFunds, t)
		assertBalance(wallet, Bitcoin(20), t)
	})
}

func assertBalance(wallet Wallet, want Bitcoin, t testing.TB) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(got, want error, t testing.TB) {
	t.Helper()
	if got == nil {
		// Fatal stops execution. Subsequent code would panic in case of nil pointer.
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		// The %q verb will print out the string in quotes
		t.Errorf("got %q, want %q", got, want)
	}
}
