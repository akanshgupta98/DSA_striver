package main

import "fmt"

/*

N = 5
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/

func main() {
	n := 6
	no := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%v ", no)
			no++
		}
		fmt.Println()
	}
}
