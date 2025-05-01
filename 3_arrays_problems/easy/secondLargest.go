package main

import (
	"fmt"
	"math"
)

func secondLargest(arr []int) int {

	max := math.MinInt
	sm := math.MinInt
	for _, val := range arr {
		if val > max {
			if sm != max {
				sm = max
			}
			max = val

		} else if val > sm && val != max {
			sm = val
		}
	}

	if sm == math.MinInt {
		sm = -1
	}
	return sm
}

func main() {

	arr := []int{1, 2, 5, 3, 4, 5, 5}

	fmt.Println("Largest element is: ", secondLargest(arr))

}

/*

[3,3,12]
M, SM
M = 3,

*/
