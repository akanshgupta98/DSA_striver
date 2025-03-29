package main

import "fmt"

func findMaxInd(start, end int, arr []int) int {
	maxInd := start

	for i := start; i <= end; i++ {
		if arr[maxInd] < arr[i] {
			maxInd = i
		}
	}
	return maxInd
}
func bubbleSort(arr []int) {

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		// maxInd := findMaxInd(0, i, arr)
		// arr[i], arr[maxInd] = arr[maxInd], arr[i]
	}
}
func main() {
	// In each iteration, the goal is that by swapping 2 numbers, at the end of iteration,
	// the maximum number will reach the end.
	arr := []int{4, 3, 11, 2, 1}
	fmt.Println("Before sorting ", arr)
	bubbleSort(arr)
	fmt.Println("After sorting ", arr)
}
