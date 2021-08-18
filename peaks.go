package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	peaks := PosPeaks{}
	peaks.Pos = []int{}
	peaks.Peaks = []int{}
	i := 1
	for i < len(array)-1 {
		if array[i] > array[i-1] {
			j := i
			for {
				if array[i] == array[i+1] {
					i++
					if i == len(array)-1 {
						return peaks
					}
				} else {
					break
				}
			}
			if array[i] > array[i+1] {
				peaks.Pos = append(peaks.Pos, j)
				peaks.Peaks = append(peaks.Peaks, array[j])
			}
		}
		i++
	}

	fmt.Println(peaks)
	return peaks
}

func main() {
	array := []int{1, 2, 3, 2, 1, 2, 2}
	fmt.Println(PickPeaks(array))
}
