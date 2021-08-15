package main

import (
	"fmt"
)

func main() {
	var data string = "idso"
	for i := range data {
		fmt.Println(data[i])
	}
}
