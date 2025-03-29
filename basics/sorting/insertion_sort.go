package main

import "fmt"

func insertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if tmp > arr[j] {
				arr[j+1] = tmp
				break

			} else {
				arr[j+1] = arr[j]
			}
		}
		if arr[0] > tmp {
			arr[0] = tmp
		}
	}
	return arr

}
func main() {
	arr := []int{10, 12, 14, 11, 8, 9, 15, 13, 1}
	fmt.Println("Before sorting", arr)
	insertionSort(arr)
	fmt.Println("After sorting", arr)

}

// The goal is that a in my sub-array, all the elements will be sorted.
// Sub-array starts with 0, and goes till end.
// If I find any element which is not in place, just shift all the elements, and then when find the correct place of new element,
// Place it there and break.

// 10,12,14,11,8,9,10,15,12
// if arr[i] > arr[i-1]
// Then sorted. Relax.
// Move ahead.
// else
//  for j=0; j<i;j++
// 		if arr[j] > arr[i]
//			shift all elements to j+1 till i-1,
//				then arr[j] = arr[i]

// 1,
// 1,2

// 1,2,4
//
