package maps

const (
	ErrNotFound   = DictionaryErr("Sorry, the word you are looking for is not in the dictionary.")
	ErrWordExists = DictionaryErr("Sorry, the word you are trying to add already exists in the dictionary.")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (entry string, searchError error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	// Type of error that Search might return
	case ErrNotFound:
		d[word] = definition
	// If search returned nil, it means that it suceeded
	case nil:
		return ErrWordExists
	// Handle any other error
	default:
		return err
	}
	// All good
	return nil
}
