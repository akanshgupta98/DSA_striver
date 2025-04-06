package main

import "fmt"

/*
N = 5
* * * * *
* * * *
* * *
* *
*
*/
func main() {
	N := 6
	fmt.Println("N = ", N)
	for i := 0; i < N; i++ {
		for j := N; j > i; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
