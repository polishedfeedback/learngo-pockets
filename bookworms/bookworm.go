package main

import (
  "os"
  "encoding/json"
)

type Bookworm struct {
  Name string `json: "name"`
  Book []Books `json: "books"`
}

type Books struct {
  Author string `json: "author"`
  Title string `json: "title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {

  // Read the file 
  f, err := os.Open(filePath)
  if err != nil{
    return nil, err
  }
  defer f.close()
  // Decode the json from the file using encoder package
  var bookworms []Bookworm
  err = json.NewDecoder(f).Decode(&bookworms)
  if err != nil {
    return nil, err
  }
  return bookworms, nil
}
