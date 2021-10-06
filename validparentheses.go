package main

import (
	"fmt"
)

/*
func count(str string, s string) int {
	strslice:=strings.Split(str, "")
	count := 0
	for _,v:= range strslice {
		if v==s {
			count++
		}
	}
	return count
}
*/

func count(str string, r rune) int {
	count := 0
	for _, v := range str {
		if v == r {
			count++
		}
	}
	return count
}

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
