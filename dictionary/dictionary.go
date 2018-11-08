package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you're looking for")

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

