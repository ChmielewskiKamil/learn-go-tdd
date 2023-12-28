package maps

import "testing"

func TestDictionarySearch(t *testing.T) {
    searchKey := "test"
    dictionary := Dictionary{searchKey: "this is just a test"}

    got := dictionary.Search(searchKey)
    want := "this is just a test"

    if got != want {
        t.Errorf("got %q want %q given, %q", got, want, searchKey)
    }
}
