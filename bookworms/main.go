package main

import (
	"fmt"
	"os"
)

func main() {
	bw, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: `err`")
		os.Exit(1)
	}
	fmt.Println(bw)
}
