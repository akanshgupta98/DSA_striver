package main

import "fmt"

/*
N=5
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
*/
func main() {
	N := 6
	fmt.Println("N = ", N)

	for i := 0; i < N; i++ {
		for j, k := N, 1; j > i; j-- {
			fmt.Printf("%d ", k)
			k++
		}
		fmt.Println()
	}

}
