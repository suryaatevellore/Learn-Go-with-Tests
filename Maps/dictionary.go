package main

// import "errors"

const (
	ErrNotFound         = DictionaryErr("Couldn't find the word you are looking for")
	ErrWordExists       = DictionaryErr("Word exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist in the dictionary")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(searchWord string) (string, error) {
	definition, ok := d[searchWord]
	if !ok {
		return "", ErrNotFound

	}
	return definition, nil
}

func (d Dictionary) Add(word, meaning string) error {
	// maps can be modified without passing them as a pointer
	// because maps are a reference type, as they hold reference to the underlying
	// data structures
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = meaning
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, meaning string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = meaning
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
