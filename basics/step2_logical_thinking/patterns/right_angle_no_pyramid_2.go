package main

import "fmt"

/*
N = 5
1
2 2
3 3 3
4 4 4 4
5 5 5 5 5
*/
func main() {
	N := 5
	for i := 1; i <= N; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}

}
