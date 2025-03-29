package main

import "fmt"

// Given N, print square * pattern. So N=6, then 6 starts horizontal, 6 starts vertical.
// N=3
// * * *
// * * *
// * * *
func main() {
	N := 6
	for row := 0; row < N; row++ {

		for col := 0; col < N; col++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}

}
