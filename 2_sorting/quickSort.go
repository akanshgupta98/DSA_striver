package main

import "fmt"

func quickSort() {

}
func main() {

	arr := []int{12, 4, 3, 2, 1, 5, 6, 9}
	fmt.Println("[Original] Arr is: ", arr)
	quickSort(arr)
	fmt.Println("[Sorted] Arr is: ", arr)
}
