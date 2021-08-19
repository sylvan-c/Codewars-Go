package main

import (
	"fmt"
	"strconv"
	"strings"
)

func LCM(nums []int) int {
	max := 0
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	flag := 0
	for i := 1; i < max; i++ {
		flag = 0
		for _, n := range nums {
			if (max*i)%n != 0 {
				flag = 1
			}
		}
		if flag == 0 {
			return i * max
		}
	}
	return 0
}

func simplifyFrac(f []int) []int {
	num := f[0]
	denom := f[1]
	for i := num; i > 0; i-- {
		if num%i == 0 && denom%i == 0 {
			num = num / i
			denom = denom / i
			i = num
		} else {
			i--
		}
		if num == 1 {
			break
		}
	}
	frac := []int{num, denom}
	return frac
}

func ConvertFracts(a [][]int) string {
	for i, f := range a {
		a[i] = simplifyFrac(f)
	}
	denoms := []int{}
	for _, f := range a {
		denoms = append(denoms, f[1])
	}
	newDenom := LCM(denoms)
	newFracs := []string{}
	for _, f := range a {
		multiplier := newDenom / f[1]
		frac := "(" + strconv.Itoa(f[0]*multiplier) + "," + strconv.Itoa(newDenom) + ")"
		newFracs = append(newFracs, frac)
	}
	fracString := strings.Join(newFracs, "")
	return fracString
}

func main() {
	fmt.Println(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}
