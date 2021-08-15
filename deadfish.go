package main

import "fmt"

func Parse(data string) []int {
	val := 0
	var out []int
	for i := range data {
		switch {
		case data[i] == 105:
			val += 1
		case data[i] == 100:
			val -= 1
		case data[i] == 115:
			val = val * val
		case data[i] == 111:
			out = append(out, val)
		}
	}
	return out
}

func main() {
	fmt.Println(Parse("ssdizbsnsdkkincuybidds"))
}
