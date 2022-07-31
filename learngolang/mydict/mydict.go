package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound       = errors.New("not Found")
	errWordExists     = errors.New("word Already Exists")
	errNoWordToUpdate = errors.New("no Word To Update")
	errNoWordToDelete = errors.New("no Word To Delete")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNoWordToUpdate
	case nil:
		d[word] = def
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNoWordToDelete
	case nil:
		delete(d, word)
	}
	return nil
}
