package main

import "fmt"

func rBubbleSort(arr []int, start, end int) {

	if start >= end {
		return
	}
	for i := start; i < end; i++ {

		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]

		}

	}
	rBubbleSort(arr, start, end-1)
}
func main() {
	arr := []int{11, 4, 5, 2, 12, 9, 10, 23, 100, 98, 76, 1}
	fmt.Println("Arr is: ", arr)
	rBubbleSort(arr, 0, len(arr)-1)
	fmt.Println("[Sorted] Arr is: ", arr)

}
