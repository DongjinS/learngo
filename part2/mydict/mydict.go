package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("NotFound")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

// Search for a word
func (d Dictionary) Search (word string) (string, error){
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}


// Add word to the dictionary
func (d Dictionary) Add (word, def string) error{
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
		return nil
	case nil:
		return errWordExists
	}
	return nil
}

// Update a given word's def in Dictionary 
func (d Dictionary) Update (word, newDef string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		d[word] = newDef
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete (word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errCantDelete
	}
	delete(d, word)
	return nil
}