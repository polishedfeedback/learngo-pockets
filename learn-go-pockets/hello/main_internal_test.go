package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T){

  type testCase struct {
    lang language
    want string
  }

  var tests = map[string] testCase {
    "English": {
      lang: "en",
      want: "Hello world",
    },

    "French": {
      lang: "fr",
      want: "Bonjour le monde",
    },
    
    "Greece": {
      lang: "el",
      want: "Χαίρετε Κόσμε",
    },
    "Japanese": {
      lang: "jp",
      want: "こんにちは世界",
    },
    "Telugu": {
      lang: "tel",
      want: "హలో వరల్డ్",
    },
   // "Empty": {
   //   lang: "",
   //   want: `Unsupported Language: ""`,
   // },
  }

  for name, tc := range tests {
    t.Run(name, func(t *testing.T){
      got := greet(tc.lang)
      if got != tc.want{
        t.Errorf("expected: %q, got %q", tc.want, got)
      }
    })
  }
}

//func TestGreet(t *testing.T) {
//	// Preparation step
//	want := "Hello world"
//
//	// Execution step
//	got := greet()
//
//	// Decision phase
//	if got != want {
//		t.Errorf("expected: %q, got %q", want, got)
//	}
//}

//func TestGreet_English(t *testing.T) {
//	lang := language("en")
//	want := "Hello world"
//
//	got := greet(lang)
//
//	if got != want {
//    // Mark this test as failed
//		t.Errorf("expected: %q, got %q", want, got)
//	}
//}
//
//func TestGreet_French(t *testing.T) {
//	lang := language("fr")
//	want := "Bonjour le monde"
//
//	got := greet(lang)
//  
//  // Using greeting here to make sure the variables are accessible while testing
//	if got != want {
//    // Mark this test as failed
//		t.Errorf("expected: %q, got %q", want, got)
//	}
//}
//
//func TestGreet_Akkadian(t *testing.T) {
//	lang := language("ak")
//	want := ""
//
//	got := greet(lang)
//  
//  // Using greeting here to make sure the variables are accessible while testing
//	if got != want {
//    // Mark this test as failed
//		t.Errorf("expected: %q, got %q", want, got)
//	}
//}
