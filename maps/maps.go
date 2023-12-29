package maps

import "errors"

var ErrNotFound = errors.New("Sorry, the word you are looking for is not in the dictionary.")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (entry string, searchError error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
    d[word] = definition
}
