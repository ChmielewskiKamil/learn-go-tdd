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

		assertError(t, err, ErrNotFound)
	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "this is a new word"
		definition := "new word"
		dictionary.Add(definition, word)

		assertDefinition(t, dictionary, definition, word)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "existing word"
		definition := "this is an existing word"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
		assertError(t, err, ErrWordExists)
	})
}

func TestDictionaryUpdate(t *testing.T) {
	existingWord := "existing word"
	existingDefinition := "this is an existing definition"

	t.Run("existing word", func(t *testing.T) {
        dictionary := Dictionary{existingWord: existingDefinition}
		newDefinition := "this is a new definition"
		dictionary.Update(existingWord, newDefinition)
		assertDefinition(t, dictionary, existingWord, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
	    dictionary := Dictionary{existingWord: existingDefinition}
		newWord := "new word"
		newDefinition := "this is a new definition"
		err := dictionary.Update(newWord, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})

	t.Run("new definition is the same as old one", func(t *testing.T) {
	    dictionary := Dictionary{existingWord: existingDefinition}
        newDefinition := "this is an existing definition"
		err := dictionary.Update(existingWord, newDefinition)

		assertError(t, err, ErrNewDefinitionIsSameAsOld)
	})

    t.Run("delete existing word", func(t *testing.T) {
        dictionary := Dictionary{existingWord: existingDefinition}
        dictionary.Delete(existingWord)

        _, err := dictionary.Search(existingWord)
        assertError(t, err, ErrNotFound)
    })
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertString(t, got, definition)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
