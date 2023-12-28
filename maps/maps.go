package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (entry string) {
    return d[word]
}
