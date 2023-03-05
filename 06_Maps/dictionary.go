package main

const (
	ErrNoWordFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists  = DictionaryErr("word already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func defaultDictionary() map[string]string {
	return map[string]string{"key0": "val0", "key1": "val1", "key2": "val2"}
}

func defaultSearch(dictionary map[string]string, key string) string {
	return dictionary[key]
}

// Custom implemantation of dictionary
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return "", ErrNoWordFound
	}
	return val, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNoWordFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNoWordFound:
		return err
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNoWordFound:
		return err
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
