package mydict

import (
	"errors"
	"fmt"
)

//dictionary type
type Dictionary map[string]string

var (
	errNotFound     = errors.New("err: not found")
	errWordExists   = errors.New("err: word exists")
	errCannotUpdate = errors.New("err: cannot update non-existing word")
	errCannotDelete = errors.New("error: cannot delete word. does not exist")
)

func (d Dictionary) Display() {
	if len(d) == 0 {
		fmt.Println("Dictionary is empty!")
	} else {
		k := 0
		for i, j := range d {
			fmt.Println(k, i, ": ", j)
			k++
		}
	}
}

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {

	//search for word, catch error
	_, err := d.Search(word)

	//check error and if not found
	//ret error other wise
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errNotFound
	}

	return nil

}

//Update a words definition
func (d Dictionary) Update(word, defin string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = defin
	case errNotFound:
		return errCannotUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCannotDelete
	}

	return nil
}
