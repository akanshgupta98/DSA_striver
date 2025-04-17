package main

import "fmt"

func main() {

	arr := []int{2, 2, 3, 4, 4, 2}

	hashMap := make(map[int]int)

	for _, val := range arr {
		hashMap[val]++
	}
	fmt.Println("Original arr is: ", arr)
	for key, val := range hashMap {
		fmt.Printf("%d occurs %d times\n", key, val)
	}
	fmt.Println()
}
