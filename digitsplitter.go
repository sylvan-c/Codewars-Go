package main

import (
	"fmt"
)

func DigitAdder(n int) int {
	remainder := 0
	sum := 0
	for n != 0 {
		remainder = n % 10
		sum += remainder
		n = n / 10
	}
	return sum
}

func DigitalRoot(n int) int {
	for n > 9 {
		n = DigitAdder(n)
	}
	return n
}

func main() {
	for {
		var n int
		fmt.Println("Enter number: ")
		fmt.Scanf("%d", &n)
		fmt.Println(DigitalRoot(n))
	}
}
