package maps

import "testing"

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

        assertString(t, got, want)
	})

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")

        if err == nil {
            t.Fatal("expected to get an error, but didn't get one")
        }

        assertError(t, err, ErrNotFound)
    })
}

func assertString(t testing.TB, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertError(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
