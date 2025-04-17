package main

import "fmt"

func mergeSort(arr []int, start, end int) {
	low := start
	high := end
	mid := low + ((high - low) / 2)
	if mid == low {
		return
	}
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
	return

}
func merge(arr []int, start, mid, end int) {
	tmp := []int{}

}

func main() {

	arr := []int{2, 1, 3, 4} //, 3, 4, 5, 15, 23, 11, 3, 4, 5}
	// low := 0
	// high := len(arr)
	// mid := (high - low) / 2
	fmt.Println("Original: ", arr)
	mergeSort(arr, 0, len(arr))
	// mergeSort(arr, mid+1, high)
	// merge(arr)

	fmt.Println(arr)

}
