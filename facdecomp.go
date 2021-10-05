package main

import (
	"fmt"
)

func isPrime(p int) bool {
	for i := 2; i < p; i++ {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int) int {
	if n == 0 || n == 1 {
		return 2
	}
	p := n + 1
	for {
		if isPrime(p) {
			return p
		}
		p++
	}
}

func Decomp(n int) []int {
	var factors []int
	p := 0
	for n > 1 {
		p = nextPrime(p)
		for {
			if n%p == 0 {
				factors = append(factors, p)
				n = n / p
				continue
			}
			break
		}
	}
	return factors
}

func main() {
	fmt.Println(Decomp(66879))
}
