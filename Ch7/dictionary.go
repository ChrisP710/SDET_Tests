package main

type Dictionary map[string]string

// var (
// 	ErrNotFound   = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, defintion string) {
	d[word] = defintion
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }
