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

func TestDictionaryAdd(t *testing.T) {
    t.Run("new word", func(t *testing.T){
        dictionary := Dictionary{}
        dictionary.Add("new word", "this is a new word")

        want := "this is a new word"
        got, err := dictionary.Search("new word")

        if err != nil {
            t.Fatal("should find added word:", err)
        }

        assertString(t, got, want)
    })

    t.Run("existing word", func(t *testing.T){
        dictionary := Dictionary{"existing word": "this is an existing word"}
        err := dictionary.Add("existing word", "this is an existing word")

        if err == nil {
            t.Fatal("should not be able to add an existing word")
        }

        assertError(t, err, ErrWordExists)
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
