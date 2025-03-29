package main

import "fmt"

func findLowestIndex(start, end int, arr []int) int {
	lowestInd := start
	for i := start; i < end; i++ {
		if arr[i] < arr[lowestInd] {
			lowestInd = i
		}
	}
	return lowestInd
}

// O(N^2)
func sort(arr []int) {
	// O(N)
	for i := 0; i < len(arr); i++ {
		//O(N)
		lowestInd := findLowestIndex(i, len(arr), arr)
		arr[i], arr[lowestInd] = arr[lowestInd], arr[i]
	}
}
func main() {
	// simple sorintg.
	// Find loweest no from (i,N), and place it in i position.
	arr := []int{13, 46, 24, 52, 20, 9}
	fmt.Println("Before sorting: ", arr)
	sort(arr)
	fmt.Println("After sorting: ", arr)

	// for i:=0;i<len(arr)
}
