package maps

const (
	ErrNotFound         = DictionaryErr("Sorry, the word you are looking for is not in the dictionary.")
	ErrWordExists       = DictionaryErr("Sorry, the word you are trying to add already exists in the dictionary.")
	ErrWordDoesNotExist = DictionaryErr("Sorry, the word you are trying to update does not exist in the dictionary.")
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

func (d Dictionary) Update(word, newDefinition string) error {
    _, err := d.Search(word)
    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        d[word] = newDefinition
    default:
        return err
    }

	return nil
}
