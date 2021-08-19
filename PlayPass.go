package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func rotN(r rune, n rune) rune {
	if r >= 'a' && r <= 'z' {
		r = (r - 97 + n) % 26
		return r + 97
	} else if r >= 'A' && r <= 'Z' {
		r = (r - 65 + n) % 26
		return r + 65
	} else if r >= '0' && r <= '9' {
		return 57 + 48 - r
	}
	return r
}

func PlayPass(s string, n int) string {
	s_tmp := []rune{}
	for _, r := range s {
		s_tmp = append(s_tmp, rotN(r, rune(n)))
	}
	for i, r := range s_tmp {
		if i%2 == 0 {
			s_tmp[i] = unicode.ToUpper(r)
			fmt.Println(strconv.QuoteRune(r))
		} else {
			s_tmp[i] = unicode.ToLower(r)
			fmt.Println(strconv.QuoteRune(r))
		}
	}
	for i, j := 0, len(s_tmp)-1; i < j; i, j = i+1, j-1 {
		s_tmp[i], s_tmp[j] = s_tmp[j], s_tmp[i]
	}
	var newStr string = string(s_tmp)
	return newStr
}

func main() {
	fmt.Println(PlayPass("BORN IN 2015!", 1))
}
