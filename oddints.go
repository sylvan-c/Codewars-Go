package main

import "fmt"

func FindOdd(seq []int) int {
	for i := range seq {
		num := seq[i]
		count := 0
		for j := range seq {
			if seq[j] == num {
				count += 1
			}
		}
		if count%2 != 0 {
			return num
		}
	}
	return 1
}

func main() {
	l := []int{1, 3, 2, 1, 3}
	fmt.Println(FindOdd(l))
}
