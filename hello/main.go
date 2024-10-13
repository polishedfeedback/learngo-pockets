package main

import (
  "fmt"
  "flag"
)

func main(){
  var lang string
  flag.StringVar(&lang, "lang", "en", "The required language, e.g., en")
  flag.Parse()

  greeting := greet(language(lang))
  fmt.Println(greeting)
}

type language string

// Create a map to store the language and greeting

var phrasebook = map[language]string {
"el": "Χαίρετε Κόσμε", // Greek
"en": "Hello, world", // English
"fr": "Bonjour, le monde", // French
"ur": " دﻧﯿﺎ ہﯿﻠﻮ ", // Urdu
"vi": "Xin chào Thế Giới", // Vietnamese
"tel": "హలో, ప్రపంచం", // Telugu
}
func greet(l language) string{
  greeting, ok := phrasebook[l]
  if !ok {
    return fmt.Sprintf("unsupported language: %q", l)
  }
  return greeting
}

