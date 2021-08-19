package main

import (
	"fmt"
	"unicode"
)

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, l := range str {
		if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(alphanumeric(""))
}
