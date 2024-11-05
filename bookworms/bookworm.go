package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json "name"`
	Books []Book `json "books"`
}

type Book struct {
	Title  string
	Author string
}

// loadBookworms reads the file and returns the list of Bookworms or an error if something goes wrong
func loadBookworms(filePath string) ([]Bookworm, error) {
	// Opening a file (for read access) will give file descriptor and an error
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(f).Decode(&bookworms)

	if err != nil {
		return nil, err
	}

	return bookworms, nil
}
