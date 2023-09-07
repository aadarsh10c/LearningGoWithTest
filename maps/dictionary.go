package maps

import (
	"errors"
	"fmt"
)

// create a dictionary type
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

// create a method for our dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	fmt.Printf("value: %q of %T \n", definition, definition)
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
