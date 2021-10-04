package main

import (
	"fmt"
	"strconv"
	"strings"
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

func OrderWeight(strng string) string {
	numList := strings.Fields(strng)
	weights := make([]int, len(numList))
	for i, n := range numList {
		intN, _ := strconv.Atoi(n)
		weights[i] = DigitAdder(intN)
	}

	return "hh"
}

func main() {
	fmt.Println(OrderWeight("11 3 444  1 5"))
}
