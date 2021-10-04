package main

import (
	"fmt"
	"sort"
)

func ValidateSolution(m [][]int) bool {
	if ValidateRows(m)+ValidateCols(m)+ValidateSquares(m) == 0 {
		return true
	}
	return false
}

func ValidateRows(m [][]int) int {
	row_tmp := make([]int, 9)
	for j := range m {
		for i, v := range m[j] {
			row_tmp[i] = v
		}
		//fmt.Println(row_tmp)
		sort.Ints(row_tmp)
		for i, v := range row_tmp {
			if v != i+1 {
				return 1
			}
		}
	}
	//fmt.Println(m)
	return 0
}

func ValidateCols(m [][]int) int {
	col_tmp := make([]int, 9)
	for j := 0; j < len(m); j++ {
		for i := 0; i < len(m); i++ {
			col_tmp[i] = m[i][j]
		}
		//fmt.Println(col_tmp)
		sort.Ints(col_tmp)
		for i, v := range col_tmp {
			if v != i+1 {
				return 1
			}
		}
	}
	//fmt.Println(m)
	return 0
}

func ValidateSquares(m [][]int) int {
	var square_tmp []int
	for i := 0; i < len(m); i += 3 {
		for j := 0; j < len(m); j += 3 {
			square_tmp = []int{}
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					square_tmp = append(square_tmp, m[i+k][j+l])
				}
			}
			//fmt.Println(square_tmp)
			sort.Ints(square_tmp)
			for i, v := range square_tmp {
				if v != i+1 {
					return 1
				}
			}
		}
	}
	return 0
}

var m = [][]int{{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9}}

func main() {
	fmt.Println(ValidateSolution(m))
}
