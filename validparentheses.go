package main

import (
	"fmt"
)

func ValidParentheses(parens string) bool {
	var t rune = 0
	for _, v := range parens {
		if v == '(' {
			t += 1
		} else {
			t -= 1
		}
		if t < 0 {
			return false
		}
	}
	if t != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(ValidParentheses("((()))"))
}
