package main

import (
  "fmt"
  "flag"
)

func main() {
  var lang string
  flag.StringVar(&lang, "lang", "en", "The required language, e.g. en")
  flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// Language represents the lanugage code
type language string

// Create a phrasebook which contains both language and greeting
var phrasebook = map[language] string{
  "el": "Χαίρετε Κόσμε",
  "en": "Hello world",
  "fr": "Bonjour le monde",
  "jp": "こんにちは世界",
  "tel": "హలో వరల్డ్",
}

// greet() returns a greeting to the world in a specified language
func greet(l language) string {
//	switch l {
//	case "en":
//		return "Hello world"
//	case "fr":
//		return "Bonjour le monde"
//	default:
//		return ""
//	}
  greeting, ok := phrasebook[l]
  if !ok {
    return fmt.Sprintf("unsupported language: %q", l)
  }
  return greeting
}
