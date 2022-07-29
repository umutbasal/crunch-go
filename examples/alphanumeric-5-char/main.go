package main

import (
	"io/ioutil"

	"github.com/umutbasal/haveibeentracked/pkg/crunch"
)

func main() {
	b, err := crunch.GenerateFromCharset(5, 5, "ualpha")
	if err != nil {
		panic(err)
	}

	// write to file
	err = ioutil.WriteFile("crunch.txt", b, 0644)
	if err != nil {
		panic(err)
	}
}
