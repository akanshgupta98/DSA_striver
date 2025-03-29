package main

import "fmt"

/*
N = 6
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
1 2 3 4 5 6
*/
func main() {
	N := 3
	for row := 0; row < N; row++ {
		for col := 1; col <= row+1; col++ {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

}
