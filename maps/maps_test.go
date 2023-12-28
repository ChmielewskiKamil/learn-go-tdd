package maps

import "testing"

func TestDictionarySearch(t *testing.T) {
    searchKey := "test"
    dictionary := map[string]string{searchKey: "this is just a test"}

    got := Search(dictionary, searchKey)
    want := "this is just a test"

    if got != want {
        t.Errorf("got %q want %q given, %q", got, want, searchKey)
    }
}
