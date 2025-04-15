package main

import "fmt"

func RevArr(arr []int, st, end int) {

	if st < end {
		arr[st], arr[end] = arr[end], arr[st]
		RevArr(arr, st+1, end-1)
	}

}
func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original arr is: ", arr)
	RevArr(arr, 0, len(arr)-1)
	fmt.Println("Rev arr is: ", arr)

}
