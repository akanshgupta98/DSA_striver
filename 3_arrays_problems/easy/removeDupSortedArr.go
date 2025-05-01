package main

import (
	"fmt"
)

func removeDup(arr []int) int {
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j++
			arr[j] = arr[i]
		}
	}
	return j + 1

}

func main() {

	arr := []int{-2, 2, 4, 4, 4, 4, 5, 5}
	r := removeDup(arr)
	fmt.Println("After removal of duplicates: ", arr)
	fmt.Println("Unique elements are: ", r)

}

/*

[3,3,12]
O/P: [3,12]


*/
