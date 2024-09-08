package main

import (
  "fmt"
  "github.com/peshamlg/learngo/dictionaryexample/mydict"
)

func main() {
  dictionary := mydict.Dictionary{}

  baseWord := "hello"
  def, errSearch := dictionary.Search(baseWord)
  if errSearch != nil {
    fmt.Println(errSearch)
  } else {
    fmt.Println(def)
  }
  
  errAdd := dictionary.Add(baseWord, "First Word")
  if errAdd != nil {
    fmt.Println(errAdd)
  } else {
    fmt.Println("The word is added to dictonary")
    def, errSearch = dictionary.Search(baseWord)
    if errSearch != nil {
      fmt.Println(errSearch)
    } else {
      fmt.Println(def)
    }
  }

  errUpdate := dictionary.Update(baseWord, "Second Word")
  if errUpdate != nil {
    fmt.Println(errUpdate)
  } else {
    fmt.Println("The definition of word is updated")
    def, errSearch = dictionary.Search(baseWord)
    if errSearch != nil {
      fmt.Println(errSearch)
    } else {
      fmt.Println(def)
    }
  }

  errDelete := dictionary.Delete(baseWord)
  if errDelete != nil {
    fmt.Println(errDelete)
  } else {
    fmt.Println("The word is deleted")
    def, errSearch = dictionary.Search(baseWord)
    if errSearch != nil {
      fmt.Println(errSearch)
    } else {
      fmt.Println(def)
    }
  }
}
