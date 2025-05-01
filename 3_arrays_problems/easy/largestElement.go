package main

import (
	"fmt"
	"math"
)

func largestElement(arr []int) int {

	max := math.MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func main() {

	arr := []int{3, 3}

	fmt.Println("Largest element is: ", largestElement(arr))

}
