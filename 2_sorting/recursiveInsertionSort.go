package main

import "fmt"

func rInsertionSort(arr []int, start, end int) {
	if end == len(arr) {
		return
	}
	i := end
	tmp := arr[i]
	for i = end; i > start; i-- {

		if arr[i-1] > tmp {
			arr[i] = arr[i-1]
		} else {
			break
		}

	}

	arr[i] = tmp

	rInsertionSort(arr, start, end+1)
}

func main() {
	arr := []int{5, 4}
	fmt.Println("[Original] Arr is: ", arr)
	rInsertionSort(arr, 0, 0)
	fmt.Println("[Sorted] Arr is: ", arr)

}
