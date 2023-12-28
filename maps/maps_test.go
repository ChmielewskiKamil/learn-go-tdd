package maps

import "testing"

func TestDictionarySearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
