package maps

import "testing"

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

    t.Run("unknown word", func(t *testing.T) {
        _, err := dictionary.Search("unknown")
        want := "Sorry, the word you are looking for is not in the dictionary."

        if err == nil {
            t.Fatal("expected to get an error, but didn't get one")
        }

        if err.Error() != want {
            t.Errorf("got %q want %q", err, want)
        }
    })
}
