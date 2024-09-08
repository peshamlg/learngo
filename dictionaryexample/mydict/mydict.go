package mydict

import "errors"

type Dictionary map[string]string

// Search a word
func (d Dictonary) Search (word string) (string, error) {
  def, exist := d[word]
  if exist {
    return def, nil
  }
  return "", errors.New("The word is not found")
}

// Add a word to the dictonary
func (d Dictonary) Add(word, def string) error {
  _, err := d.Search(word)
  if err == nil {
    d[word] = def
  } else {
    return errors.New("The word is already exist")
  }
  return nil
}

// Update a word
func (d Dictonary) Update(word, def string) error {
  _, err := d.Search(word)
  if err != nil {
    d[word] = def
  } else {
    return errors.New("The word is not exist")
  }
  return nil
}

// Delete a word
func (d Dictonary) Delete (word string) error {
  _, err := d.Search(word)
  if err != nil {
    delete(d, word)
  } else {
    return errors.New("The word is not exist")
  }
  return nil
}
