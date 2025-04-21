package main

import "fmt"

func mergeSort(arr []int, start, end int) {
	low := start
	high := end
	mid := (low + high) / 2
	if low >= high {
		return
	}
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
	return

}
func merge(arr []int, start, mid, end int) {
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		if arr[i] < arr[j] {
			tmp = append(tmp, arr[i])
			// tmp[k] = arr[i]
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
		// k++
	}

	for i <= mid {
		tmp = append(tmp, arr[i])
		// k++
		i++
	}

	for j <= end {
		tmp = append(tmp, arr[j])
		j++
	}

	for i = start; i <= end; i++ {
		arr[i] = tmp[i-start]

	}

}

func main() {

	arr := []int{2, 1, 3, 4, 3, 4, 5, 15, 23, 11, 3, 4, 5}
	fmt.Println("Original: ", arr)
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
