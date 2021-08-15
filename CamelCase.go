package main

import (
	"fmt"
	"strings"
)

func CamelCase(s string) string {
	s = strings.Title(s)
	fmt.Println(s)
	s = strings.Join(strings.Split(s, " "), "")
	return s
}

func main() {
	fmt.Println(CamelCase("hello my name is jehovah"))
}
