package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func Anagrams(word string, words []string) []string {
	res := []string{}
	sortedWord := SortString(word)
	for _, w := range words {
		if sortedWord == SortString(w) {
			res = append(res, w)
		}
	}
	return res
}

func main() {
	fmt.Println(Anagrams("hello", []string{"aaa", "bbb", "olleh", "hello"}))
}
