package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) [2]string {
	st1 := strings.SplitAfter(st, "")
	st2 := strings.SplitAfter(st, "")
	for i := 0; i < len(st1); i += 2 {
		st1[i] = strings.ToUpper(st1[i])
	}
	for i := 1; i < len(st2); i += 2 {
		st2[i] = strings.ToUpper(st2[i])
	}
	l := [2]string{strings.Join(st1, ""), strings.Join(st2, "")}
	return l
}

func main() {
	fmt.Println(Capitalize("abcdef"))
}
