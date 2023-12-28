package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (entry string, searchError error) {
    if d[word] == "" {
        return "", errors.New("Sorry, the word you are looking for is not in the dictionary.")
    }
    return d[word], nil 
}
