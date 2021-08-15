package main

import (
	"fmt"
)

func triangle(n int) int {
	return n*(n+1)/2
}

func firstInRow(n int) int {
	return triangle(n-1)+1
}

func RowSumOddNumbers (n int) int {
	firstNum:=firstInRow(n)*2-1
	sum := firstNum*n + 2*(triangle(n-1))
	return sum
}

func main() {
	for {
		var n int
		fmt.Println("Enter number: ")
		fmt.Scanf("%d", &n)
		fmt.Println(RowSumOddNumbers(n))
	}
}