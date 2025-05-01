package main

import (
	"fmt"
)

func checkAsc(arr []int) (result bool) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return
		}
	}

	return true
}

func main() {

	arr := []int{1, 3, 3, 1}

	fmt.Println("Largest element is: ", checkAsc(arr))

}
