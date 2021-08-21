package main

import (
	"fmt"
	"math"
)

func Going(n int) float64 {
	var next float64 = 1
	var series float64 = 0
	for i := 0; i < n; i++ {
		series += float64(next)
		next = next * (float64(1) / (float64(n) - float64(i)))
	}
	return math.Floor((series * 1000000)) / 1000000
}

func main() {
	for i := 1; i < 100; i++ {
		fmt.Println(Going(i))
	}
}
