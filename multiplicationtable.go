package main

import (
	"fmt"
)

func MultiplicationTable(size int) [][]int {
	var table = make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			table[i] = append(table[i], (i+1)*(j+1))
		}
	}
	return table
}

func main() {
	fmt.Println(MultiplicationTable(5))
}
