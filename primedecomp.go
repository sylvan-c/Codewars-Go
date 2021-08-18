package main

import (
	"fmt"
	"strconv"
	"strings"
)

var primes = []int{2}

func NextPrime(p int) int {
	i := p + 1
	for {
		flag := 0
		flag = 0
		for _, j := range primes {
			if i%j == 0 {
				flag = 1
			}
		}
		if flag == 0 {
			primes = append(primes, i)
			return i
		}
		i++
	}
}

func PrimeFactors(n int) string {
	facList := []string{}
	count := 0
	p := 2
	for {
		var str string
		count = 0
		powstr := ""
		for {
			if n%p == 0 {
				n = n / p
				count += 1
			} else {
				if count == 0 {
					break
				} else {
					if count > 1 {
						powstr = "**" + strconv.Itoa(count)
					}
				}
				str = "(" + strconv.Itoa(p) + powstr + ")"
				break
			}
		}
		facList = append(facList, str)
		if n == 1 {
			break
		}
		p = NextPrime(p)
	}
	factors := strings.Join(facList, "")
	return factors
}

func main() {
	fmt.Println(PrimeFactors(7775460))
}
