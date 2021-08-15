package main

import (
	"fmt"
)

func findMax(list [3]int) int {
	max:=0
	for i:=0; i<len(list); i++ {
		if list[i]>max {
			max=list[i]
		}
	}
	return max
}

func IsTriangle(a,b,c int) bool {
	var list = [3]int{a,b,c}
	max:=0
	for i:=0; i<3; i++ {
		if list[i]>max {
			max=list[i]
		}
	}
	if a+b+c>2*max {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsTriangle(5,6,2))
}