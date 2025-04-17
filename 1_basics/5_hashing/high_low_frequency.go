package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 2, 3, 4, 4, 2}

	hashMap := make(map[int]int)

	for _, val := range arr {
		hashMap[val]++
	}
	highest, hF := math.MinInt, math.MinInt
	lowest, lF := math.MaxInt, math.MaxInt

	for key, val := range hashMap {
		if highest < val {
			highest = val
			hF = key
		}
		if lowest > val {
			lowest = val
			lF = key
		}
	}

	fmt.Printf("HIghest: %d Lowest: %d ", hF, lF)
	fmt.Println()
}
